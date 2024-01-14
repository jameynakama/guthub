package main

import (
	"flag"
	"fmt"
)

const DEFAULT_LIMIT = 25

type cfg struct {
	repoLimit int
}

func main() {
	repoLimit := flag.Int("l", DEFAULT_LIMIT, "limit of repositories to scrape")
	flag.Parse()

	cfg := cfg{
		repoLimit: *repoLimit,
	}

	run(cfg)
}

func run(cfg cfg) {
	var sh scrapeHelper

	sh.getTrendingRepos(cfg, "https://github.com/trending/")

	for _, repo := range sh.toScrape {
		// TODO: Use GH API to get text files
		// TODO: Maybe even get comments eventually
		fmt.Println(repo)
	}
}
