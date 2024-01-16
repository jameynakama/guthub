# GutHub

GutHub is a tool for grabbing the READMEs of the top 25 "trending repos" on GH.

The purpose is to save these READMEs for later scanning, to check for typos and
grammatical mistakes that can then be PRed against.

## Future plans

- Improved testing (currently missing a few lines of coverage, and no
  integration tests)
- Flag for choosing specific output directory
- Flag for choosing to write to a temp dir and temp files (and open them)

## Installation

```bash
git clone git@github.com:jameynakama/guthub
cd guthub
go install ./...
```

You can also just run the tool from the repo:

`go run ./cmd/guthub -h`

## Usage

You must have a valid GH token with read permissions for public repos. Set this
as an env var, or provide it directly in the call to the tool.

The tool will grab all top 25 repos without any args.

It will save the READMEs in `CWD/guthub-output/` with the pattern
`<owner>--<repo-name>.md`.

`$ GH_TOKEN=<your-token> guthub -l 3`

```text
Usage of guthub:
 -l int
   limit of repositories to scrape (default 25)
```

Example: `GH_TOKEN=<your-token> guthub -l 3`

This will fetch the READMEs of the first 3 repos on the trending page.

## License

This project is licensed under the terms of the MIT license.
