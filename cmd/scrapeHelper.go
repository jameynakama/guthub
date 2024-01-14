package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type scrapeHelper struct {
	toScrape []string
}

func (s *scrapeHelper) getTrendingRepos(cfg cfg, url string) {
	c := colly.NewCollector()

	repoSelector, found := os.LookupEnv("REPO_SELECTOR")
	if !found {
		repoSelector = "h2.h3.lh-condensed > a[href]"
	}

	c.OnHTML(repoSelector, func(e *colly.HTMLElement) {
		if len(s.toScrape) >= cfg.repoLimit {
			return
		}
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			link = e.Request.AbsoluteURL(link)
		}
		s.toScrape = append(s.toScrape, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Getting repos to scrape...")
	})

	c.Visit(url)
}
