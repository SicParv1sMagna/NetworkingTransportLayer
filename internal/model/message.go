package model

import "time"

type Message struct {
	StringMessage string    `json:"string_message"`
	SenderName    string    `json:"sender_name"`
	Time          time.Time `json:"time"`
}
