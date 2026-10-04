package dto

import (
	"mime/multipart"
	"time"
)

type EventResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	OrganizerId int       `json:"organizer_id"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Speakers    []string  `json:"speakers"`
	Categories  []string  `json:"categories"`
	CommunityId int       `json:"community_id"`
	Attendees   int       `json:"attendees"`
}

type CreateEventRequest struct {
	ID          int                  `form:"id"`
	OrganizerId int                  `form:"organizer_id"`
	Title       string               `form:"title"`
	Image       multipart.FileHeader `form:"image"`
	Description string               `form:"description"`
	Location    string               `form:"location"`
	Capacity    int                  `form:"capacity"`
	StartTime   time.Time            `form:"start_time"`
	EndTime     time.Time            `form:"end_time"`
	Categories  []int                `form:"categories"`
	Speakers    []int                `form:"speakers"`
	NewSpeakers string               `form:"new_speakers"`
	CommunityId int                  `form:"community_id"`
}
