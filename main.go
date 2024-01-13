package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type cfg struct {
	repoLimit int
}

type scrapeHelper struct {
	toScrape []string
	scraped  []string
}

func (s *scrapeHelper) getTrendingRepos(cfg cfg) {
	c := colly.NewCollector()

	c.OnHTML("h2.h3.lh-condensed > a[href]", func(e *colly.HTMLElement) {
		if len(s.toScrape) >= cfg.repoLimit {
			return
		}
		link := e.Attr("href")
		s.toScrape = append(s.toScrape, e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Getting repos to scrape...")
	})

	c.Visit("https://github.com/trending")
}

func (s *scrapeHelper) scrapeReadme(url string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Scraping %s...\n", r.URL.String())
	})

	c.OnHTML("span.author > a[href]", func(e *colly.HTMLElement) {
		fmt.Println("\tAuthor: " + strings.TrimSpace(e.Text))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("\tFinished scraping " + r.Request.URL.String())
		s.scraped = append(s.scraped, r.Request.URL.String())
	})

	c.Visit(url)
}

func main() {
	repoLimit := flag.Int("l", 25, "limit of repositories to scrape")
	flag.Parse()

	cfg := cfg{
		repoLimit: *repoLimit,
	}

	run(cfg)
}

func run(cfg cfg) {
	var sh scrapeHelper

	sh.getTrendingRepos(cfg)

	for _, repo := range sh.toScrape {
		sh.scrapeReadme(repo)
	}
}
