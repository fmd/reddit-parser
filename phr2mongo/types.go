package main

type Comment struct {
	Text      string `json:"text"`
	Points    string `json:"points"`
	PointsInt int    `json:"points_int"`
	User      string `json:"user"`
	Time      string `json:"time"`
	Post      string `json:"post"`
}

type Post struct {
	Title     string    `json:"title"`
	User      string    `json:"user"`
	Time      string    `json:"time"`
	Points    string    `json:"points"`
	PointsInt int       `json:"points_int"`
	Id        string    `json:"id"`
	Comments  []Comment `json:"comments,omitempty"`
}

type Sub struct {
	Posts []Post `json:"posts"`
}
