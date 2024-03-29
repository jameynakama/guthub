package main

import (
	"testing"

	"github.com/jameynakama/assert"
	"github.com/jameynakama/guthub/cmd/githubapi"
	"github.com/jameynakama/guthub/cmd/guthub/testhelpers"
)

type testLogger struct{}

func (l *testLogger) Info(v ...any)  {}
func (l *testLogger) Error(v ...any) {}
func (l *testLogger) Debug(v ...any) {}

func TestGetRepos(t *testing.T) {
	server, err := testhelpers.NewTestServer("trending.html")
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	expRepos := []githubapi.Repo{
		{Owner: "someone", Name: "ugly-cats"},
		{Owner: "someone", Name: "a-spicy-a-meat-a-ball-a"},
		{Owner: "someone", Name: "the-sisters-karamazov"},
	}

	sh := newScrapeHelper(&testLogger{})
	sh.getTrendingRepos(server.URL, 3)

	assert.Equal(t, sh.repos, expRepos)
}

func TestGetRepoDefaultLimit(t *testing.T) {
	server, err := testhelpers.NewTestServer("trending.html")
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	expRepos := []githubapi.Repo{
		{Owner: "someone", Name: "ugly-cats"},
		{Owner: "someone", Name: "a-spicy-a-meat-a-ball-a"},
		{Owner: "someone", Name: "the-sisters-karamazov"},
		{Owner: "someone", Name: "poop-jokes"},
		{Owner: "someone", Name: "cycle-jordan"},
	}

	sh := newScrapeHelper(&testLogger{})
	sh.getTrendingRepos(server.URL, DEFAULT_LIMIT)

	assert.Equal(t, sh.repos, expRepos)
}
