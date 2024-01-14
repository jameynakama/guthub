package githubapi

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/go-github/v58/github"
	"github.com/jameynakama/guthub/cmd/logging"
	"golang.org/x/oauth2"
)

type Repo struct {
	Author string
	Name   string
}

type GitHubAPI interface {
	GetReadme(
		ctx context.Context,
		owner,
		repo string,
		opts *github.RepositoryContentGetOptions,
	) (
		file *github.RepositoryContent,
		directoryContent []*github.RepositoryContent,
		resp *github.Response, err error,
	)
}

type GitHubAPIClient struct {
	client    *github.Client
	authToken string
	logger    logging.Logger
}

func NewGitHubAPIClient(authToken string, l logging.Logger) *GitHubAPIClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &GitHubAPIClient{
		client:    github.NewClient(tc),
		authToken: authToken,
		logger:    l,
	}
}

func (c *GitHubAPIClient) GetReadmes(repos []Repo) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.authToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	var wg sync.WaitGroup
	errCh := make(chan error, len(repos))
	readmeCh := make(chan string, len(repos))

	for _, repo := range repos {
		// TODO: Use GH API to get text files
		// TODO: Maybe even get comments eventually
		wg.Add(1)

		go func(repo Repo) {
			defer wg.Done()

			c.logger.Info(fmt.Sprintf("Fetching %q README", repo.Name))

			readme, err := getReadme(ctx, client, repo)
			if err != nil {
				errCh <- err
				return
			}

			readmeCh <- readme
		}(repo)

		go func() {
			wg.Wait()
			close(errCh)
			close(readmeCh)
		}()

		for err := range errCh {
			c.logger.Error(err)
		}

		for readme := range readmeCh {
			c.logger.Info(readme)
		}
	}
}

func getReadme(ctx context.Context, client *github.Client, repo Repo) (string, error) {
	readme, _, err := client.Repositories.GetReadme(ctx, repo.Author, repo.Name, nil)
	if err != nil {
		return "", err
	}

	content, err := readme.GetContent()
	if err != nil {
		return "", err
	}

	return content, nil
}
