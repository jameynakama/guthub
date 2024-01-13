package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	repoLimit := flag.Int("l", 25, "limit of repositories to scrape")
	flag.Parse()

	var reposVisited int

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnHTML("h2.h3.lh-condensed > a[href]", func(e *colly.HTMLElement) {
		reposVisited++
		if reposVisited > *repoLimit {
			return
		}

		fmt.Printf("Will scrape %s\n", e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s...\n-----\n", r.URL)
	})

	run(c)
}

func run(c *colly.Collector) {
	c.Visit("https://github.com/trending")
}
