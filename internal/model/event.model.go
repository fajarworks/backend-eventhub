package model

import "time"

type Event struct {
	ID          int        `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Image       string     `db:"image"`
	Location    string     `db:"location"`
	Capacity    int        `db:"capacity"`
	StartTime   time.Time  `db:"start_time"`
	EndTime     time.Time  `db:"end_time"`
	Categories  []string   `db:"categories"`
	Speakers    []string   `db:"speakers"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	CommunityId int        `db:"community_id"`
	OrganizerId int        `db:"organizer_id"`
}
