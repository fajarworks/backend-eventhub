package model

type AdminOverview struct {
	TotalUser      int `db:"total_user"`
	TotalEvent     int `db:"total_event"`
	TotalCommunity int `db:"total_community"`
	AvgFillRate    int `db:"avg_fill_rate"`
}
