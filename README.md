Parsing Reddit with ParseHub
============================

## Introduction

[ParseHub](http://www.parsehub.com) is a great tool for parsing reddit. My ParseHub projects for parsing Reddit are included in this repository.

## Prerequisites

* MongoDB
* Go 1.0+

## Background

ParseHub gets us as far as having a `posts` JSON object, containing several `post` objects, which in turn contain several `comment` objects.
* The `phr2mongo` package imports ParseHub data into MongoDB collections.
* The `words` package parses post titles and comments and creates collections detailing word counts in the post titles and comments.
* The `phrase` package matches a regex phrase against comments and returns matching comments.
* The `points` package parses comments and creates a collection where each appearance of a word is multiplied by the number of upvotes received by the comment. 

## Installation

```
go get github.com/fmd/reddit-parser
```
