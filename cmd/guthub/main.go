package main

import (
	"flag"
	"fmt"
	"os"
)

const DEFAULT_LIMIT = 25

type cfg struct {
	repoLimit int
	url       string
	infoLog   Logger
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
		infoLog:   infoLog,
	}

	run(cfg)
}

func run(cfg cfg) {
	sh := newScrapeHelper(cfg.infoLog)
	sh.getTrendingRepos(cfg.url, cfg.repoLimit)

	for _, repo := range sh.repos {
		// TODO: Use GH API to get text files
		// TODO: Maybe even get comments eventually
		cfg.infoLog.Info(fmt.Sprintf("%+v", repo))
	}
}
