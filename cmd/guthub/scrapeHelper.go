package main

import (
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type scrapeHelper struct {
	toScrape []string
	logger   Logger
}

func newScrapeHelper(logger Logger) *scrapeHelper {
	return &scrapeHelper{
		logger: logger,
	}
}

func (s *scrapeHelper) getTrendingRepos(url string, limit int) {
	c := colly.NewCollector()

	repoSelector, found := os.LookupEnv("REPO_SELECTOR")
	if !found {
		repoSelector = "h2.h3.lh-condensed > a[href]"
	}

	c.OnHTML(repoSelector, func(e *colly.HTMLElement) {
		if len(s.toScrape) >= limit {
			return
		}
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			link = e.Request.AbsoluteURL(link)
		}
		s.toScrape = append(s.toScrape, link)
	})

	c.OnRequest(func(r *colly.Request) {
		s.logger.Info("Getting repos to scrape...")
	})

	c.Visit(url)
}
