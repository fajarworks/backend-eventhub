package service

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type EventService struct {
	repo *repository.EventRepo
}

func NewEventService(repo *repository.EventRepo) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (s *EventService) GetEvents(ctx context.Context, categoryId int, location, search, sortBy string, page, limit int) ([]dto.EventResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 6
	}

	filter := repository.EventFilter{
		CategoryId: categoryId,
		Location:   location,
		Search:     search,
		SortBy:     sortBy,
		Limit:      limit,
		Offset:     (page - 1) * limit,
	}

	events, err := s.repo.GetEvents(ctx, filter)
	if err != nil {
		return nil, err
	}

	result := make([]dto.EventResponse, 0, len(events))
	for _, e := range events {
		categories, err := s.repo.GetCategoriesByEventId(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		attendees, err := s.repo.GetEventAttendees(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, dto.EventResponse{
			ID:         e.ID,
			Title:      e.Title,
			Image:      e.Image,
			Location:   e.Location,
			Capacity:   e.Capacity,
			StartTime:  e.StartTime,
			Categories: categories,
			Attendees:  attendees,
		})
	}
	return result, nil
}
