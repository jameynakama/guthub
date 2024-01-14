package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/google/go-github/v58/github"
	"golang.org/x/oauth2"
)

const DEFAULT_LIMIT = 25

type cfg struct {
	repoLimit int
	url       string
	logger    Logger
}

type repo struct {
	author string
	name   string
}

func main() {
	repoLimit := flag.Int("l", DEFAULT_LIMIT, "limit of repositories to scrape")
	flag.Parse()

	infoLog := NewGutHubLogger(os.Stdout, os.Stderr, "[GUTHUB] ", 0)

	cfg := cfg{
		repoLimit: *repoLimit,
		url:       "https://github.com/trending/",
		logger:    infoLog,
	}

	run(cfg)
}

func run(cfg cfg) {
	sh := newScrapeHelper(cfg.logger)
	sh.getTrendingRepos(cfg.url, cfg.repoLimit)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	var wg sync.WaitGroup
	errCh := make(chan error, len(sh.repos))
	readmeCh := make(chan string, len(sh.repos))

	for _, r := range sh.repos {
		// TODO: Use GH API to get text files
		// TODO: Maybe even get comments eventually
		wg.Add(1)

		go func(thing repo) {
			defer wg.Done()

			cfg.logger.Info(fmt.Sprintf("Fetching %q README", thing.name))

			readme, err := getReadme(ctx, client, thing)
			if err != nil {
				errCh <- err
				return
			}

			readmeCh <- readme
		}(r)

		go func() {
			wg.Wait()
			close(errCh)
			close(readmeCh)
		}()

		for err := range errCh {
			cfg.logger.Error(err)
		}

		for readme := range readmeCh {
			cfg.logger.Info(readme)
		}
	}
}

func getReadme(ctx context.Context, client *github.Client, repo repo) (string, error) {
	readme, _, err := client.Repositories.GetReadme(ctx, repo.author, repo.name, nil)
	if err != nil {
		return "", err
	}

	content, err := readme.GetContent()
	if err != nil {
		return "", err
	}

	return content, nil
}
