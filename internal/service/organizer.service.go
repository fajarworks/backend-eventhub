package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizerService struct {
	repo *repository.OrganzierRepo
	db   *pgxpool.Pool
}

func NewOrganizerService(repo *repository.OrganzierRepo, db *pgxpool.Pool) *OrganizerService {
	return &OrganizerService{
		repo: repo,
		db:   db,
	}

}

func (s *OrganizerService) GetOrganizerstats(ctx context.Context, organizerId int) (dto.OrganizerStats, error) {
	stats, err := s.repo.GetOrganizerstats(ctx, s.db, organizerId)
	if err != nil {
		return dto.OrganizerStats{}, nil
	}
	return dto.OrganizerStats{
		TotalEvents:    stats.TotalEvents,
		TotalAttendees: stats.TotalAttandees,
		AvgFillRate:    stats.AvgFillRate,
	}, nil
}

func (s *OrganizerService) CreateEvent(ctx context.Context, organizerId int, imagePath string, req dto.CreateEventRequest, newSpeaker []dto.Speaker) (int, error) {
	if req.StartTime.After(req.EndTime) && !req.EndTime.IsZero() {
		return 0, apperror.ErrInvalidTimeRange
	}
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	event := model.Event{
		Title:       req.Title,
		Description: req.Description,
		Image:       imagePath,
		Location:    req.Location,
		Capacity:    req.Capacity,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		CommunityId: req.CommunityId,
		OrganizerId: organizerId,
	}
	id, err := s.repo.CreateEvent(ctx, tx, event)
	if err != nil {
		return 0, err
	}

	if err := s.repo.AddEventSpeakers(ctx, tx, id, req.Speakers); err != nil {
		return 0, err
	}
	for _, speaker := range newSpeaker {
		if err := s.repo.AddNewEventSpeakers(ctx, tx, id, speaker.Name, speaker.JobPosition); err != nil {
			return 0, err
		}
	}

	if err := s.repo.AddEventCategories(ctx, tx, id, req.Categories); err != nil {
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}
	return id, nil

}
