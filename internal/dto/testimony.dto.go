package dto

import "time"

type Testimonials struct {
	ID           int       `json:"id"`
	Fullname     string    `json:"fullname"`
	PhotoProfile string    `json:"photo_profile"`
	JobPosition  string    `json:"job_position"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"created_at"`
}
