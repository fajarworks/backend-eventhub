package model

import "time"

type User struct {
	Id           int        `db:"id"`
	Role         string     `db:"role"`
	Fullname     string     `db:"fullname"`
	Email        string     `db:"email"`
	Password     string     `db:"password"`
	PhotoProfile *string    `db:"photo_profile"`
	Bio          *string    `db:"bio"`
	JobPosition  *string    `db:"job_position"`
	Location     *string    `db:"location"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}
