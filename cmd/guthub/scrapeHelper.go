package main

import (
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/jameynakama/guthub/cmd/githubapi"
	"github.com/jameynakama/guthub/cmd/logging"
)

type scrapeHelper struct {
	repos  []githubapi.Repo
	logger logging.Logger
}

func newScrapeHelper(l logging.Logger) *scrapeHelper {
	return &scrapeHelper{
		logger: l,
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
		s.repos = append(s.repos, githubapi.Repo{
			Author: linkParts[0],
			Name:   linkParts[1],
		})
	})

	c.OnRequest(func(r *colly.Request) {
		s.logger.Info("Getting repos to scrape...")
	})

	c.Visit(url)
}
