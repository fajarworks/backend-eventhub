package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type OrganizerService struct {
	repo *repository.OrganzierRepo
}

func NewOrganizerService(repo *repository.OrganzierRepo) *OrganizerService {
	return &OrganizerService{
		repo: repo,
	}

}

func (s *OrganizerService) GetOrganizerstats(ctx context.Context, organizerId int) (dto.OrganizerStats, error) {
	stats, err := s.repo.GetOrganizerstats(ctx, organizerId)
	if err != nil {
		return dto.OrganizerStats{}, nil
	}
	return dto.OrganizerStats{
		TotalEvents:    stats.TotalEvents,
		TotalAttendees: stats.TotalAttandees,
		AvgFillRate:    stats.AvgFillRate,
	}, nil
}
