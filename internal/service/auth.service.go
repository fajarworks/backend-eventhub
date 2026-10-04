package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/fajarworks/backend-eventhub/internal/pkg"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	repo *repository.AuthRepo
	rdb  *redis.Client
}

func NewAuthService(repo *repository.AuthRepo, rdb *redis.Client) *AuthService {
	return &AuthService{
		repo: repo,
		rdb:  rdb,
	}
}

func (s *AuthService) NewUser(ctx context.Context, body dto.RegisterRequest) error {

	if body.Email == "" || len(body.Password) == 0 {
		return apperror.ErrEmptyField
	}

	if err := pkg.ValidateLengthPass(body.Password); err != nil {
		return err
	}

	if !pkg.ValidateEmailFormat(body.Email) {
		return apperror.ErrInvalidEmailFormat
	}
	_, err := s.repo.FindUserByEmail(ctx, body.Email)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if err == nil {
		return apperror.ErrAlreadyExist
	}

	hashConf := pkg.NewRecomHashConfig()
	hashedPass, _ := hashConf.GenHash(body.Password)

	if err := s.repo.CreateNewUser(ctx, model.User{
		Fullname: body.FullName,
		Email:    body.Email,
		Password: hashedPass,
	}); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) User(ctx context.Context, body dto.LoginRequest) (string, error) {
	if body.Email == "" || len(body.Password) == 0 {
		return "", apperror.ErrEmptyField
	}

	if err := pkg.ValidateLengthPass(body.Password); err != nil {
		return "", err
	}

	user, err := s.repo.FindUserByEmail(ctx, body.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {

			return "", apperror.ErrWrongEmailPass
		}
		return "", err
	}
	if err := pkg.ComparePassAndHash(body.Password, user.Password); err != nil {
		return "", apperror.ErrWrongEmailPass
	}

	claims := pkg.NewJWTClaims(user.Id, user.Role)

	return claims.GenToken()
}

func (s *AuthService) Logout(ctx context.Context, userId int, jti string, expiresAt time.Time) error {

	key := fmt.Sprintf("eventhub:blacklist:%s", jti)
	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		return nil
	}

	if err := s.rdb.Set(ctx, key, userId, ttl).Err(); err != nil {
		return err
	}
	return nil
}
