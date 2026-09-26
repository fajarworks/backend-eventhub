package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepo
}

func NewCategoryService(repo *repository.CategoryRepo) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s CategoryService) GetCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dto.CategoryResponse, 0, len(categories))

	for _, cat := range categories {
		result = append(result, dto.CategoryResponse{
			ID:   cat.ID,
			Name: cat.Name,
		})
	}
	return result, nil
}
