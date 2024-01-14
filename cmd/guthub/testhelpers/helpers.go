package testhelpers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

var errLog = log.New(os.Stderr, "[ERROR] ", 0)

func NewTestServer(filename string) (*httptest.Server, error) {
	var retErr error

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			html, err := os.ReadFile(filepath.Join("testdata", filename))
			if err != nil {
				retErr = err
			}
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintln(w, string(html))
		}),
	)
	return server, retErr
}
