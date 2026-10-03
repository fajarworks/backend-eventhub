package model

import "time"

type Community struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type CommunityMember struct {
	ID           int    `db:"id"`
	FullName     string `db:"fullname"`
	PhotoProfile string `db:"photo_profile"`
	JobPosition  string `db:"job_position"`
}

type PopularCommunities struct {
	ID             int       `db:"id"`
	Name           string    `db:"name"`
	Image          string    `db:"image"`
	Categories     []string  `db:"categories"`
	UpcomingEvents int       `db:"upcoming_event"`
	Members        int       `db:"members"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
