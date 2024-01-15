package githubapi

// import (
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/google/go-github/v58/github"
// 	"github.com/jameynakama/assert"
// )

// type mockGitHubAPIClient struct{}

// func (c *mockGitHubAPIClient) GetReadme(
// 	ctx context.Context,
// 	owner,
// 	repo string,
// 	opts *github.RepositoryContentGetOptions,
// ) (
// 	file *github.RepositoryContent,
// 	resp *github.Response,
// 	err error,
// ) {
// 	mockReadme := &github.RepositoryContent{Content: github.String("mock readme")}
// 	mockResponse := &github.Response{
// 		Response: &http.Response{
// 			StatusCode: http.StatusOK,
// 		},
// 	}
// 	return mockReadme, mockResponse, nil
// }

// func createMockServer() *httptest.Server {
// 	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		switch r.URL.Path {
// 		case "/repos/bike-tyson/is-even/readme":
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte(`[
// 				{
// 					"sha": "1234567890",
// 					"commit": {
// 						"message": "commit message"
// 					}
// 				}
// 			]`))
// 		default:
// 			w.WriteHeader(http.StatusNotFound)
// 		}
// 	}))
// }

// func TestGetReadme(t *testing.T) {
// 	mockClient := &mockGitHubAPIClient{}
// 	// mockLogger := &mockLogger{}
// 	apiClient := NewGitHubAPIClient("mock token", nil)
// 	apiClient.client = mockClient

// 	repos := []Repo{
// 		{Author: "bike-tyson", Name: "is-even"},
// 		{Author: "cycle-jordan", Name: "is-odd"},
// 	}

// 	apiClient.GetReadmes(repos)

// 	assert.Equal(t, repos[0].Readme, "mock readme")
// }
