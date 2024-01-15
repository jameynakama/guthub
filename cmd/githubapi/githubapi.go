package githubapi

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/go-github/v58/github"
	"github.com/jameynakama/guthub/cmd/logging"
)

type Repo struct {
	Owner  string
	Name   string
	Readme string
}

type RepositoriesClient interface {
	GetReadme(
		ctx context.Context,
		owner,
		repo string,
		opts *github.RepositoryContentGetOptions,
	) (*github.RepositoryContent, *github.Response, error)
}

type GutHubHelper struct {
	ctx        context.Context
	RepoClient RepositoriesClient
	logger     logging.Logger
}

func NewGutHubClient(ctx context.Context, rClient RepositoriesClient, l logging.Logger) *GutHubHelper {
	return &GutHubHelper{
		ctx:        ctx,
		RepoClient: rClient,
		logger:     l,
	}
}

func (c *GutHubHelper) GetReadmes(repos []Repo, outputDir string) {
	var wg sync.WaitGroup
	errCh := make(chan error, len(repos))
	repoCh := make(chan Repo, len(repos))

	for _, repo := range repos {
		// TODO: Use GH API to get text files
		// TODO: Maybe even get comments eventually
		wg.Add(1)

		go func(repo Repo) {
			defer wg.Done()

			c.logger.Info(fmt.Sprintf("Fetching %q README", repo.Name))

			readme, _, err := c.RepoClient.GetReadme(c.ctx, repo.Owner, repo.Name, nil)
			if err != nil {
				errCh <- err
				return
			}

			content, err := readme.GetContent()
			if err != nil {
				errCh <- err
				return
			}

			repo.Readme = content

			repoCh <- repo
		}(repo)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(repoCh)
	}()

	for err := range errCh {
		c.logger.Error(err)
	}

	for repo := range repoCh {
		filename := fmt.Sprintf("%s--%s.md", repo.Owner, repo.Name)
		if err := writeReadmeToFile(repo, outputDir, filename); err != nil {
			c.logger.Error(err)
		} else {
			c.logger.Info(fmt.Sprintf("Wrote %q README to file at %s/%s", repo.Name, outputDir, filename))
		}
	}
}

func writeReadmeToFile(repo Repo, dirName, filename string) error {
	// TODO: Allow user to specify output directory
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dirName, filename), []byte(repo.Readme), 0644); err != nil {
		return err
	}

	return nil
}
