package models

type User struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Karma int `json:"karma"`
}

type Users []User
