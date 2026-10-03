package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type TestimonyService struct {
	repo *repository.TestimonyRepo
}

func NewTestimonyService(repo *repository.TestimonyRepo) *TestimonyService {
	return &TestimonyService{
		repo: repo,
	}
}

func (s *TestimonyService) GetTestimonials(ctx context.Context) ([]dto.Testimonials, error) {
	testimonials, err := s.repo.GetTestimonials(ctx)
	if err != nil {
		return nil, err
	}

	testimony := make([]dto.Testimonials, 0, len(testimonials))
	for _, v := range testimonials {
		testimony = append(testimony, dto.Testimonials{
			ID:           v.ID,
			Fullname:     v.Fullname,
			PhotoProfile: v.PhotoProfile,
			JobPosition:  v.JobPosition,
			Message:      v.Message,
			CreatedAt:    v.CreatedAt,
		})
	}
	return testimony, nil
}
