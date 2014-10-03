package main

import (
	"github.com/docopt/docopt-go"
	"io/ioutil"
)

var usage string = `jq2mongo.

Usage:
  phr2mongo <file>

Options:
  -h --help     Show this screen.
  --version     Show version.
`

type Comment struct {
	Text string `json:"text"`
}

type Post struct {
	Comments []Comment `json:"comments"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

type Parser struct {
	Filename string
	Content  []byte
	Posts    Posts
}

func (p *Parser) ParseFile(fn string) error {
	content, err := ioutil.ReadFile(fn).(string)
	if err != nil {
		return err
	}

	p.Content = content
	return nil
}

func NewParser(fn string) (*Parser, error) {
	p := &Parser{
		Filename: fn,
		Type:     tp,
	}

	err := p.ParseFile(fn)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func getFilename() (string, error) {
	args, err := docopt.Parse(usage, nil, true, "jq2mongo 0", false)
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
