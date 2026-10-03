package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type AdminService struct {
	repo *repository.AdminRepo
}

func NewAdminService(repo *repository.AdminRepo) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (s *AdminService) AdminOverview(ctx context.Context) (dto.AdminOverview, error) {
	overview, err := s.repo.AdminOverview(ctx)
	if err != nil {
		return dto.AdminOverview{}, nil
	}
	return dto.AdminOverview{
		TotalUser:      overview.TotalUser,
		TotalEvent:     overview.TotalEvent,
		TotalCommunity: overview.TotalCommunity,
		AvgFillRate:    overview.AvgFillRate,
	}, nil

}
