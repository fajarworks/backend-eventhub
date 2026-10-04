package model

type Speaker struct {
	Name        string `db:"name"`
	JobPosition string `db:"job_position"`
}
