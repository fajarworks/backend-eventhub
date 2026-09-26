package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserRepo(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetDetailUser(ctx context.Context, id int) (model.User, error) {
	return s.repo.FindUserById(ctx, id)
}
