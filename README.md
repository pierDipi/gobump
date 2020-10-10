# gobump

`gobump` changes the version of any `go.mod` file in a repository to a specified Go version.

## Installation

`go get github.com/pierdipi/gobump/cmd/gobump`

## Usage

```bash
Usage of gobump:
  -exclude-regex string
    	Exclude regex
  -target string
    	Go target version
      
Example:
  gobump --target 1.15 --exclude-regex ^third_party
```
