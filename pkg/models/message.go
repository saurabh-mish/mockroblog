package models

import "time"

type Message struct {
	Id int `json:"-"`
	UserFrom string `json:"user_from"`
	UserTo string `json:"user_to"`
	MessageContent string `json:"message_content"`
	Flag bool `json:"flag"`
	MessagedOn time.Time `json:"messaged_on"`
}
