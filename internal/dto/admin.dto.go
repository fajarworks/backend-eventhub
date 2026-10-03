package dto

type AdminOverview struct {
	TotalUser      int `json:"total_user"`
	TotalEvent     int `json:"total_event"`
	TotalCommunity int `json:"total_community"`
	AvgFillRate    int `json:"avg_fill_rate"`
}
