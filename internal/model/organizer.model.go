package model

type OrganizerStats struct {
	TotalEvents    int `db:"total_events"`
	TotalAttandees int `db:"total_attendees"`
	AvgFillRate    int `db:"avg_fill_rate"`
}
