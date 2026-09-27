package dto

import "time"

type EventResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Image      string    `json:"image"`
	Location   string    `json:"location"`
	Capacity   int       `json:"capacity"`
	StartTime  time.Time `json:"start_time"`
	Categories []string  `json:"categories"`
	Attendees  int       `json:"attendees"`
}
