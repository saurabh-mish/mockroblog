package models

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	PostContent string `json:"post_content"`
	Community string `json:"community"`
}

type Posts []Post
