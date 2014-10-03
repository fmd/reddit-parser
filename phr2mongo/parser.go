package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Parser struct {
	Filename string
	Content  []byte
	Sub      Sub
}

// Unmarshal parses the JSON from []bytes content into the *Parsers `Sub` object.
func (p *Parser) Unmarshal() error {
	var err error

	// Unmarshal the JSON from the file into the `Sub` object.
	err = json.Unmarshal(p.Content, &p.Sub)
	if err != nil {
		return err
	}

	// Loop over the posts in the sub,
	for pI := range p.Sub.Posts {
		post := p.Pub.Posts[pI]

		// Convert the points into integers.
		ip, err := strconv.Atoi(post.Points)
		if err == nil {
			post.PointsInt = ip
		}

		// Loop over the comments in the post,
		for cI := range post.Comments {
			comment := post.Comments[cI]

			// Convert the comments into integers.
			ip, err := strconv.Atoi(comment.Points)
			if err != nil {
				comment.PointsInt = ip
			}

			// Assign the post's ID to the comment.
			comment.Post = post.Id
		}
	}

	return nil
}

// Parse parses the file into the Parser.
func (p *Parser) Parse(fn string) error {

	// Read the file into content.
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	// Assign content to the Parser.
	p.Content = content
	return nil
}

// NewParser returns a parser, complete with a filled in `Sub` object with parsed JSON data.
func NewParser(fn string) (*Parser, error) {
	var err error

	// Create the struct pointer instance.
	p := &Parser{
		Filename: fn,
	}

	// Parse the file.
	err = p.Parse(fn)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON in the file.
	err = p.Unmarshal()
	if err != nil {
		return nil, err
	}

	// Return the *Parser.
	return p, nil
}
