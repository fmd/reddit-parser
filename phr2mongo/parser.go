package main

import (
	"io/ioutil"
)

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
