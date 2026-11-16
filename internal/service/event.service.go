package service

import (
	"context"
	"log"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/sse"
)

type EventService struct {
	repo   *repository.EventRepo
	broker *sse.Broker
}

func NewEventService(repo *repository.EventRepo, broker *sse.Broker) *EventService {
	return &EventService{
		repo:   repo,
		broker: broker,
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

func (s *EventService) GetEventAttendees(ctx context.Context, eventId int) (int, error) {
	attendees, err := s.repo.GetEventAttendees(ctx, eventId)
	if err != nil {
		return 0, err
	}
	return attendees, nil
}

func (s *EventService) JoinEvent(ctx context.Context, userId, eventId int) error {
	if err := s.JoinEvent(ctx, userId, eventId); err != nil {
		return err
	}
	count, err := s.repo.GetEventAttendees(ctx, eventId)
	if err != nil {
		log.Printf("failed to get attendee count : %v", err)
		return err
	}

	s.broker.Publish(eventId, count)
	return nil
}

func (s *EventService) LeaveEvent(ctx context.Context, userId, eventId int) error {
	if err := s.repo.LeaveEvent(ctx, userId, eventId); err != nil {
		return err
	}

	count, err := s.repo.GetEventAttendees(ctx, eventId)
	if err != nil {
		log.Printf("failed to get attendee count : %v", err)
		return err
	}
	s.broker.Publish(eventId, count)
	return nil
}

func (s *EventService) GetDetail(ctx context.Context, eventId int) (dto.EventResponse, error) {
	event, err := s.repo.GetDetailEvent(ctx, eventId)

	if err != nil {
		return dto.EventResponse{}, err
	}

	attendees, err := s.repo.GetEventAttendees(ctx, eventId)
	if err != nil {
		return dto.EventResponse{}, err
	}

	categories, err := s.repo.GetCategoriesByEventId(ctx, eventId)
	if err != nil {
		return dto.EventResponse{}, err
	}

	speakers, err := s.repo.GetEventSpeakers(ctx, eventId)

	return dto.EventResponse{
		ID:         eventId,
		Title:      event.Title,
		Image:      event.Image,
		Location:   event.Location,
		Capacity:   event.Capacity,
		Speakers:   speakers,
		StartTime:  event.StartTime,
		Categories: categories,
		Attendees:  attendees,
	}, nil

}

func (s *EventService) GetUpcomingEvents(ctx context.Context) ([]dto.EventResponse, error) {
	events, err := s.repo.GetUpcomingEvent(ctx)

	if err != nil {
		return nil, err
	}
	data := make([]dto.EventResponse, 0, len(events))

	for _, e := range events {
		categories, err := s.repo.GetCategoriesByEventId(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		attendees, err := s.repo.GetEventAttendees(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		data = append(data, dto.EventResponse{
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
	return data, nil
}

func (s *EventService) GetEventsByUserId(ctx context.Context, userId int) ([]dto.EventResponse, error) {
	events, err := s.repo.GetEventsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	data := make([]dto.EventResponse, 0, len(events))

	for _, e := range events {
		categories, err := s.repo.GetCategoriesByEventId(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		attendees, err := s.repo.GetEventAttendees(ctx, e.ID)
		if err != nil {
			return nil, err
		}
		data = append(data, dto.EventResponse{
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
	return data, nil
}

func (s *EventService) SaveEvent(ctx context.Context, userId, eventId int) error {
	err := s.repo.SaveEvent(ctx, userId, eventId)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) RemoveSavedEvent(ctx context.Context, userId, eventId int) error {
	err := s.repo.RemoveSavedEvent(ctx, userId, eventId)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetJoinedEventIds(ctx context.Context, userId int) ([]int, error) {
	return s.repo.GetJoinedEventIds(ctx, userId)
}

func (s *EventService) GetSavedEventIds(ctx context.Context, userId int) ([]int, error) {
	return s.repo.GetSavedEventIds(ctx, userId)
}
