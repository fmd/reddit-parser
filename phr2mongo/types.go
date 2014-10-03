package main

type Comment struct {
	Text string `json:"text"`
}

type Post struct {
	Comments []Comment `json:"comments"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}
