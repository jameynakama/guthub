package main

import (
	"fmt"
	"testing"

	"github.com/jameynakama/assert"
	"github.com/jameynakama/guthub/cmd/guthub/testhelpers"
)

type testLogger struct{}

func (l *testLogger) Info(v ...any)  {}
func (l *testLogger) Error(v ...any) {}

func TestGetRepos(t *testing.T) {
	server, err := testhelpers.NewTestServer("trending.html")
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	expUrls := []string{
		"https://example.com/ugly-cats/",
		"https://example.com/a-spicy-a-meat-a-ball-a/",
		"https://example.com/the-sisters-karamazov/",
	}

	sh := newScrapeHelper(&testLogger{})
	sh.getTrendingRepos(server.URL, 3)

	assert.Equal(t, sh.toScrape, expUrls)
}

func TestGetRepoDefaultLimit(t *testing.T) {
	server, err := testhelpers.NewTestServer("trending.html")
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	expUrls := []string{
		"https://example.com/ugly-cats/",
		"https://example.com/a-spicy-a-meat-a-ball-a/",
		"https://example.com/the-sisters-karamazov/",
		"https://example.com/poop-jokes/",
		"https://example.com/cycle-jordan/",
	}

	sh := newScrapeHelper(&testLogger{})
	sh.getTrendingRepos(server.URL, DEFAULT_LIMIT)

	assert.Equal(t, sh.toScrape, expUrls)
}

func TestGetReposRelativeURLs(t *testing.T) {
	server, err := testhelpers.NewTestServer("trending_relative_urls.html")
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	expUrls := []string{
		fmt.Sprintf("%s/ugly-cats/", server.URL),
	}

	sh := newScrapeHelper(&testLogger{})
	sh.getTrendingRepos(server.URL, 1)

	assert.Equal(t, sh.toScrape, expUrls)
}
