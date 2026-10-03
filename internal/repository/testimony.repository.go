package repository

import (
	"context"
	"log"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TestimonyRepo struct {
	db *pgxpool.Pool
}

func NewTestimonyRepo(db *pgxpool.Pool) *TestimonyRepo {
	return &TestimonyRepo{
		db: db,
	}

}

func (t *TestimonyRepo) GetTestimonials(ctx context.Context) ([]model.Testimonials, error) {
	sql := `SELECT t.id, u.fullname, COALESCE(u.photo_profile, '') AS photo_profile, u.job_position, t.message, t.created_at
		FROM testimonials t
		JOIN users u ON u.id = t.user_id 
		ORDER BY t.created_at ASC`

	rows, err := t.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	var testimonials []model.Testimonials
	defer rows.Close()
	for rows.Next() {
		var testi model.Testimonials
		if err := rows.Scan(&testi.ID, &testi.Fullname, &testi.PhotoProfile, &testi.JobPosition, &testi.Message, &testi.CreatedAt); err != nil {
			return nil, err
		}
		testimonials = append(testimonials, testi)
	}
	log.Println("testimony", testimonials)
	return testimonials, nil
}
