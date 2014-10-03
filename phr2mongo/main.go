package main

import (
	"github.com/docopt/docopt-go"
)

var usage string = `phr2mongo.

Usage:
  phr2mongo <mongohost> <mongodb> <filename>

Options:
  -h --help     Show this screen.
  --version     Show version.
`

type Options struct {
	MongoHostname string
	MongoDatabase string
	Filename      string
}

func parseArgs() (*Options, error) {
	args, err := docopt.Parse(usage, nil, true, "phr2mongo 0", false)
	if err != nil {
		return nil, err
	}

	return &Options{
		MongoHostname: args["<mongohost>"].(string),
		MongoDatabase: args["<mongodb"].(string),
		Filename:      args["<filename>"].(string),
	}, nil
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		panic(err)
	}

	parser, err := NewParser(opts.Filename)
	if err != nil {
		panic(err)
	}

	mongo, err := NewMongo(opts.MongoHostname, opts.MongoDatabase)
	if err != nil {
		panic(err)
	}
}
