package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/go-github/v58/github"
	"github.com/jameynakama/guthub/cmd/githubapi"
	"github.com/jameynakama/guthub/cmd/logging"
	"golang.org/x/oauth2"
)

const DEFAULT_LIMIT = 25

type cfg struct {
	repoLimit int
	url       string
	logger    logging.Logger
	ghClient  *githubapi.GutHubHelper
}

func main() {
	repoLimit := flag.Int("l", DEFAULT_LIMIT, "limit of repositories to scrape")
	flag.Parse()

	logger := logging.NewGutHubLogger(os.Stdout, os.Stdout, os.Stderr, "[GUTHUB] ", 0)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)
	client := githubapi.NewGutHubClient(ctx, ghClient.Repositories, logger)

	cfg := cfg{
		repoLimit: *repoLimit,
		url:       "https://github.com/trending/",
		logger:    logger,
		ghClient:  client,
	}

	run(cfg)
}

func run(cfg cfg) {
	sh := newScrapeHelper(cfg.logger)
	sh.getTrendingRepos(cfg.url, cfg.repoLimit)

	cfg.ghClient.GetReadmes(sh.repos, "guthub-output")
}
