package repository

import (
	"context"

	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT id, role, password FROM users WHERE email =$1"
	args := []any{email}
	var data model.User
	if err := r.db.QueryRow(ctx, sql, args...).Scan(&data.Id, &data.Role, &data.Password); err != nil {
		return model.User{}, err
	}
	return data, nil

}

func (r *AuthRepo) CreateNewUser(ctx context.Context, data model.User) error {
	sql := "INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3)"
	args := []any{data.Fullname, data.Email, data.Password}

	cmd, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return apperror.ErrNoRowsAffected
	}
	return nil

}
