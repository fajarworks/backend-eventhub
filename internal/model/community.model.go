package model

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
