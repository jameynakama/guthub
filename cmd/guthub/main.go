package main

import (
	"flag"
	"fmt"
)

const DEFAULT_LIMIT = 25

type cfg struct {
	repoLimit int
	url       string
}

func main() {
	repoLimit := flag.Int("l", DEFAULT_LIMIT, "limit of repositories to scrape")
	flag.Parse()

	cfg := cfg{
		repoLimit: *repoLimit,
		url:       "https://github.com/trending/",
	}

	run(cfg)
}

func run(cfg cfg) {
	var sh scrapeHelper

	sh.getTrendingRepos(cfg.url, cfg.repoLimit)

	for _, repo := range sh.toScrape {
		// TODO: Use GH API to get text files
		// TODO: Maybe even get comments eventually
		fmt.Println(repo)
	}
}
