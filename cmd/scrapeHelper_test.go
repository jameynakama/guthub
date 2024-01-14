package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/jameynakama/assert"
)

func newTestServer(filename string) *httptest.Server {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			html, err := os.ReadFile(filepath.Join("testdata", filename))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading testdata/%s: %v\n", filename, err)
				os.Exit(1)
			}
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintln(w, string(html))
		}),
	)
	return server
}

func TestGetRepos(t *testing.T) {
	server := newTestServer("trending.html")
	defer server.Close()

	expUrls := []string{
		"https://example.com/ugly-cats/",
		"https://example.com/a-spicy-a-meat-a-ball-a/",
		"https://example.com/the-sisters-karamazov/",
	}

	var sh scrapeHelper
	sh.getTrendingRepos(cfg{repoLimit: 3}, server.URL)

	assert.Equal(t, sh.toScrape, expUrls)
}

func TestGetRepoDefaultLimit(t *testing.T) {
	server := newTestServer("trending.html")
	defer server.Close()

	expUrls := []string{
		"https://example.com/ugly-cats/",
		"https://example.com/a-spicy-a-meat-a-ball-a/",
		"https://example.com/the-sisters-karamazov/",
		"https://example.com/poop-jokes/",
		"https://example.com/cycle-jordan/",
	}

	var sh scrapeHelper
	sh.getTrendingRepos(cfg{repoLimit: DEFAULT_LIMIT}, server.URL)

	assert.Equal(t, sh.toScrape, expUrls)
}

func TestGetReposRelativeURLs(t *testing.T) {
	server := newTestServer("trending_relative_urls.html")
	defer server.Close()

	expUrls := []string{
		fmt.Sprintf("%s/ugly-cats/", server.URL),
	}

	var sh scrapeHelper
	sh.getTrendingRepos(cfg{repoLimit: 1}, server.URL)

	assert.Equal(t, sh.toScrape, expUrls)
}
