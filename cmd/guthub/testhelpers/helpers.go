package testhelpers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func NewTestServer(filename string) *httptest.Server {
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

func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "", 0)
}
