package model

import "time"

type Segment struct {
	ID            time.Time `json:"id"`
	TotalSegments uint      `json:"total_segments"`
	SenderName    string    `json:"sender_name"`
	SegmentNumber uint      `json:"segment_number"`
	Payload       string    `json:"payload"`
	HadError      bool      `json:"had_error"`
}
