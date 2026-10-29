package dto

import "time"

type CommunityResponse struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Categories     []string `json:"categories"`
	Members        int      `json:"members"`
	UpcomingEvents int      `json:"upcoming_events"`
}

type CommunityMembers struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
	JobPosition  string `json:"job_position"`
}

type PopularCommunities struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Image          string    `json:"image"`
	Categories     []string  `json:"categories"`
	UpcomingEvents int       `json:"upcoming_event"`
	Members        int       `json:"members"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserCommunityResponse struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Categories     []string `json:"categories"`
	Members        int      `json:"members"`
	UpcomingEvents int      `json:"upcoming_event"`
}

type CommunitiesResponse struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Categories     []string `json:"categories"`
	Members        int      `json:"members"`
	UpcomingEvents int      `json:"upcoming_event"`
}

type CommunitiesFilter struct {
	Search     string
	CategoryID *int
	Page       int
	Limit      int
}

type CommunityPageMetaData struct {
	Next      string `json:"next"`
	Prev      string `json:"prev"`
	TotalData int    `json:"total_data"`
}
