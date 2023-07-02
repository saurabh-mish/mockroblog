package models

type VoteAction struct {
	Upvote int `json:"upvote"`
	Downvote int `json:"downvote"`
}
