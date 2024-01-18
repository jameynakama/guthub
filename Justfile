alias t := test
alias tf := test-force
alias tc := test-cover
alias tch := test-cover-html
alias tcf := test-cover-func
alias b := build

test:
  go test -v ./...

test-force:
  go test --count=1 -v ./...

test-cover TYPE:
  go test -coverprofile cover.out ./... && go tool cover -{{TYPE}} cover.out

test-cover-html:
  go test -coverprofile cover.out ./... && go tool cover -html cover.out

test-cover-func:
  go test -coverprofile cover.out ./... && go tool cover -func cover.out

build VERSION:
  go build -v -ldflags="-X 'main.Version=v{{VERSION}}'" -o bin/guthub ./cmd/guthub
