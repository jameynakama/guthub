package main

import (
	"flag"
	"os"

	"github.com/jameynakama/guthub/cmd/githubapi"
	"github.com/jameynakama/guthub/cmd/logging"
)

const DEFAULT_LIMIT = 25

type cfg struct {
	repoLimit int
	url       string
	logger    logging.Logger
	ghClient  *githubapi.GitHubAPIClient
}

func main() {
	repoLimit := flag.Int("l", DEFAULT_LIMIT, "limit of repositories to scrape")
	flag.Parse()

	logger := logging.NewGutHubLogger(os.Stdout, os.Stderr, "[GUTHUB] ", 0)
	client := githubapi.NewGitHubAPIClient(os.Getenv("GH_TOKEN"), logger)

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

	cfg.ghClient.GetReadmes(sh.repos)
}
