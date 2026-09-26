package service

import (
	"context"
	"errors"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/fajarworks/backend-eventhub/internal/pkg"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	repo *repository.AuthRepo
}

func NewAuthService(repo *repository.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
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
