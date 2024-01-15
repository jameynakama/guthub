package githubapi

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-github/v58/github"
	"github.com/jameynakama/assert"
)

type mockRepositoriesClient struct{}

type mockRepoContent struct {
	Content string
}

func (m *mockRepositoriesClient) GetReadme(
	ctx context.Context,
	owner,
	repo string,
	opts *github.RepositoryContentGetOptions,
) (*github.RepositoryContent, *github.Response, error) {
	content := "# Hello"
	mockContent := &github.RepositoryContent{
		Content: &content,
	}
	if strings.HasPrefix(owner, "ERROR") {
		return nil, nil, fmt.Errorf(owner)
	}
	return mockContent, nil, nil
}

type mockLogger struct {
	errorLog []string
}

func (m *mockLogger) Info(v ...any)  {}
func (m *mockLogger) Debug(v ...any) {}
func (m *mockLogger) Error(v ...any) {
	m.errorLog = append(m.errorLog, fmt.Sprintln(v...))
}

func TestGetReadmes(t *testing.T) {
	mReposClient := &mockRepositoriesClient{}
	mGuthubHelper := NewGutHubClient(context.Background(), mReposClient, &mockLogger{})

	tempOutDir := t.TempDir()
	repos := []Repo{
		{
			Owner: "bike-tyson",
			Name:  "is-even",
		},
		{
			Owner: "cycle-jordan",
			Name:  "is-odd",
		},
	}

	mGuthubHelper.GetReadmes(repos, tempOutDir)

	for _, repo := range repos {
		file, err := os.ReadFile(filepath.Join(tempOutDir, fmt.Sprintf("%s--%s.md", repo.Owner, repo.Name)))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "# Hello", string(file))
	}
}

func TestGetReadmesError(t *testing.T) {
	mLogger := &mockLogger{}
	mReposClient := &mockRepositoriesClient{}
	mGuthubHelper := NewGutHubClient(context.Background(), mReposClient, mLogger)

	tempOutDir := t.TempDir()
	repos := []Repo{
		{Owner: "ERROR bike-tyson", Name: "is-even"},
		{Owner: "ERROR cycle-jordan", Name: "is-odd"},
		{Owner: "chill-smith", Name: "is-true"},
	}

	mGuthubHelper.GetReadmes(repos, tempOutDir)

	assert.Equal(t, mLogger.errorLog[0], "ERROR bike-tyson\n")
	assert.Equal(t, mLogger.errorLog[1], "ERROR cycle-jordan\n")

	files, err := os.ReadDir(tempOutDir)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(files), 1)
}
