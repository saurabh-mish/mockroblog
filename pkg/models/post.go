package models

import "time"

type Post struct {
	Id *int `json:"-"`
	Title string `json:"title"`
	PostContent string `json:"post_content"`
	Community string `json:"community"`
	Username string `json:"username"`
	Url string `json:"url"`
	PostedOn time.Time `json:"posted_on"`
	Upvotest int `json:"upvotes"`
	Downvotes int `json:"downvotes"`
}

type Posts []Post
