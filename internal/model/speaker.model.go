package model

type Speaker struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	JobPosition string `db:"job_position"`
}
