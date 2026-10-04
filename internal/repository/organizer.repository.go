package repository

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}
type OrganzierRepo struct {
}

func NewOrganizerRepo() *OrganzierRepo {
	return &OrganzierRepo{}
}

func (r *OrganzierRepo) GetOrganizerstats(ctx context.Context, db DBTX, organizerId int) (model.OrganizerStats, error) {
	sql := `SELECT COUNT(DISTINCT e.id) AS total_events, 
		COUNT(DISTINCT ue.user_id) AS total_attendees, 
		COALESCE(
			(COUNT(ue.user_id)::numeric / NULLIF(SUM(e.capacity) ,0)) * 100
		,0) AS avg_fill_rate
		FROM events e
		LEFT JOIN user_event ue ON ue.event_id = e.id
		WHERE e.organizer_id = $1`
	var stats model.OrganizerStats
	if err := db.QueryRow(ctx, sql, organizerId).Scan(&stats.TotalEvents, &stats.TotalAttandees, &stats.AvgFillRate); err != nil {
		return model.OrganizerStats{}, err
	}
	return stats, nil

}

func (r *OrganzierRepo) CreateEvent(ctx context.Context, db DBTX, event model.Event) (int, error) {
	sql := `
		INSERT INTO events (title, description, image, location, capacity, start_time, end_time, community_id, organizer_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	args := []any{
		event.Title,
		event.Description,
		event.Image,
		event.Location,
		event.Capacity,
		event.StartTime,
		event.EndTime,
		event.CommunityId,
		event.OrganizerId,
	}

	var id int
	if err := db.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrganzierRepo) AddEventSpeakers(ctx context.Context, db DBTX, eventId int, speakerIds []int) error {
	if len(speakerIds) == 0 {
		return nil
	}

	for _, speakerId := range speakerIds {
		sql := `INSERT INTO event_speaker (event_id, speaker_id) VALUES ($1, $2)`
		if _, err := db.Exec(ctx, sql, eventId, speakerId); err != nil {
			return err
		}
	}

	return nil
}

func (r *OrganzierRepo) AddEventCategories(ctx context.Context, db DBTX, eventId int, categoryIds []int) error {
	if len(categoryIds) == 0 {
		return nil
	}

	for _, categoryId := range categoryIds {
		sql := `INSERT INTO event_category (event_id, category_id) VALUES ($1, $2)`
		if _, err := db.Exec(ctx, sql, eventId, categoryId); err != nil {
			return err
		}
	}

	return nil
}

func (r *OrganzierRepo) AddNewEventSpeakers(ctx context.Context, db DBTX, eventId int, name, jobPosition string) error {
	sql := `INSERT INTO speakers (name, job_position) VALUES ($1, $2) returning id`
	var speakerId int

	if err := db.QueryRow(ctx, sql, name, jobPosition).Scan(&speakerId); err != nil {
		return err
	}
	sql2 := `INSERT INTO event_speaker (event_id, speaker_id) VALUES ($1, $2)`
	_, err := db.Exec(ctx, sql2, eventId, speakerId)
	if err != nil {
		return err
	}

	return nil

}
