package repository

import (
	"context"

	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
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

func (r *UserRepo) UpdateProfile(ctx context.Context, userId int, photoProfile, fullName, location, jobPosition, bio *string) (model.User, error) {
	sql := `UPDATE users 
	SET photo_profile = COALESCE ($1, photo_profile) ,
	fullname = COALESCE($2, fullname), 
	location = COALESCE ($3, location), job_position = COALESCE($4, job_position),
	bio = COALESCE ($5, bio),
	updated_at = NOW()
	WHERE id = $6
	RETURNING id, photo_profile, fullname, location, bio, updated_at`

	var user model.User
	err := r.db.QueryRow(ctx, sql, photoProfile, fullName, location, jobPosition, bio, userId).Scan(&user.Id, &user.PhotoProfile, &user.Fullname, &user.Location, &user.Bio, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) ChangePassword(ctx context.Context, userId int, password string) error {
	sql := `UPDATE users
	set password = $1
	WHERE id = $2`
	args := []any{password, userId}
	cmdTag, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return apperror.ErrNoRowsAffected
	}
	return nil
}
