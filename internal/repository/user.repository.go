package repository

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindUserById(ctx context.Context, id int) (model.User, error) {
	sql := "SELECT id, role, email,fullname , password, photo_profile, location, bio, job_position, created_at FROM users WHERE id=$1"
	args := []any{id}
	var data model.User
	if err := r.db.QueryRow(ctx, sql, args...).Scan(
		&data.Id,
		&data.Role,
		&data.Email,
		&data.Fullname,
		&data.Password,
		&data.PhotoProfile,
		&data.Location,
		&data.Bio,
		&data.JobPosition,
		&data.CreatedAt); err != nil {
		return data, err
	}
	return data, nil
}
