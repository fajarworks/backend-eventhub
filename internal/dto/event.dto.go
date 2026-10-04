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
	Title       string               `form:"title" binding:"required"`
	Image       multipart.FileHeader `form:"image" binding:"required"`
	Description string               `form:"description"`
	Location    string               `form:"location" binding:"required"`
	Capacity    int                  `form:"capacity" binding:"required,gt=0"`
	StartTime   time.Time            `form:"start_time" binding:"required"`
	EndTime     time.Time            `form:"end_time" binding:"required"`
	Categories  []int                `form:"categories"`
	Speakers    []int                `form:"speakers"`
	NewSpeakers string               `form:"new_speakers"`
	CommunityId int                  `form:"community_id"`
}
