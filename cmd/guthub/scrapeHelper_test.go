package main

import (
	"fmt"
	"testing"

	"github.com/jameynakama/assert"
	"github.com/jameynakama/guthub/cmd/guthub/testhelpers"
)

func TestGetRepos(t *testing.T) {
	server := testhelpers.NewTestServer("trending.html")
	defer server.Close()

	expUrls := []string{
		"https://example.com/ugly-cats/",
		"https://example.com/a-spicy-a-meat-a-ball-a/",
		"https://example.com/the-sisters-karamazov/",
	}

	var sh scrapeHelper
	sh.getTrendingRepos(server.URL, 3)

	assert.Equal(t, sh.toScrape, expUrls)
}

func TestGetRepoDefaultLimit(t *testing.T) {
	server := testhelpers.NewTestServer("trending.html")
	defer server.Close()

	expUrls := []string{
		"https://example.com/ugly-cats/",
		"https://example.com/a-spicy-a-meat-a-ball-a/",
		"https://example.com/the-sisters-karamazov/",
		"https://example.com/poop-jokes/",
		"https://example.com/cycle-jordan/",
	}

	var sh scrapeHelper
	sh.getTrendingRepos(server.URL, DEFAULT_LIMIT)

	assert.Equal(t, sh.toScrape, expUrls)
}

func TestGetReposRelativeURLs(t *testing.T) {
	server := testhelpers.NewTestServer("trending_relative_urls.html")
	defer server.Close()

	expUrls := []string{
		fmt.Sprintf("%s/ugly-cats/", server.URL),
	}

	var sh scrapeHelper
	sh.getTrendingRepos(server.URL, 1)

	assert.Equal(t, sh.toScrape, expUrls)
}
