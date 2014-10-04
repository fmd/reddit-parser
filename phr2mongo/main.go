package main

import (
	"github.com/docopt/docopt-go"
)

var usage string = `phr2mongo.

Usage:
  phr2mongo import <mongohost> <mongodb> <filename>
  phr2mongo drop   <mongohost> <mongodb>

Options:
  -h --help     Show this screen.
  --version     Show version.
`

type Options struct {
	MongoHostname string
	MongoDatabase string
	Filename      string
	Action        string
}

var opts   *Options
var mongo  *Mongo
var parser *Parser

func parseArgs() (*Options, error) {
	args, err := docopt.Parse(usage, nil, true, "phr2mongo 0", false)
	if err != nil {
		return nil, err
	}

	o := &Options{
		MongoHostname: args["<mongohost>"].(string),
		MongoDatabase: args["<mongodb>"].(string),
	}

	drop   := args["drop"].(bool)
	imprt := args["import"].(bool)
	
	if drop {
		o.Action = "drop"
	} else if imprt {
		o.Action = "import"
		f := args["<filename>"].(string)
		o.Filename = f
	}

	return o, nil
}

func drop() {
	err := mongo.Drop()
	if err != nil {
		panic(err)
	}
}

func imprt() {
	parser, err := NewParser(opts.Filename)
	if err != nil {
		panic(err)
	}

	err = mongo.Import(parser.Sub)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	opts, err = parseArgs()
	if err != nil {
		panic(err)
	}

	mongo, err = NewMongo(opts.MongoHostname, opts.MongoDatabase)
	if err != nil {
		panic(err)
	}

	switch opts.Action {
		case "import":
			imprt()
			break

		case "drop":
			drop()
			break
	}
}
