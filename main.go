package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnHTML("h2.h3.lh-condensed > a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s...\n-----\n", r.URL)
	})

	run(c)
}

func run(c *colly.Collector) {
	c.Visit("https://github.com/trending")
}
