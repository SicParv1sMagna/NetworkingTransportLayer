package model

import "time"

type CollectedMessage struct {
	Message    string    `json:"message"`
	SenderName string    `json:"sender"`
	Time       time.Time `json:"timestamp"`
	Error      bool      `json:"error"`
}
