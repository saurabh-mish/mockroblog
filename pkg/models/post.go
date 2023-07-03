package models

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Community string `json:"community"`
}

type Posts []Post
