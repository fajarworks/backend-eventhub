package model

import "time"

type Testimonials struct {
	ID           int       `db:"id"`
	Fullname     string    `db:"fullname"`
	PhotoProfile string    `db:"photo_profile"`
	JobPosition  string    `db:"job_position"`
	Message      string    `db:"message"`
	CreatedAt    time.Time `db:"created_at"`
}
