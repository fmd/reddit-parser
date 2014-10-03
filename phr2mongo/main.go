package main

import (
	"github.com/docopt/docopt-go"
)

var usage string = `phr2mongo.

Usage:
  phr2mongo <file>

Options:
  -h --help     Show this screen.
  --version     Show version.
`

func getFilename() (string, error) {
	args, err := docopt.Parse(usage, nil, true, "phr2mongo 0", false)
	if err != nil {
		return "", err
	}

	return args["<file>"].(string)
}

func main() {
	f, err := getFilename()
	if err != nil {
		panic(err)
	}

	parser, err := NewParser(fn)
	if err != nil {
		panic(err)
	}
}
