package main

import (
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type scrapeHelper struct {
	repos  []repo
	logger Logger
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
		if len(s.repos) >= limit {
			return
		}
		link := e.Attr("href")
		linkParts := strings.Split(strings.Trim(link, "/"), "/")
		s.repos = append(s.repos, repo{
			author: linkParts[0],
			name:   linkParts[1],
		})
	})

	c.OnRequest(func(r *colly.Request) {
		s.logger.Info("Getting repos to scrape...")
	})

	c.Visit(url)
}
