package dto

type OrganizerStats struct {
	TotalEvents    int `json:"total_events"`
	TotalAttendees int `json:"total_attendees"`
	AvgFillRate    int `json:"avg_fill_rate"`
}
