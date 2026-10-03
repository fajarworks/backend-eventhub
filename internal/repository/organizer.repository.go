package repository

import (
	"context"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganzierRepo struct {
	db *pgxpool.Pool
}

func NewOrganizerRepo(db *pgxpool.Pool) *OrganzierRepo {
	return &OrganzierRepo{
		db: db,
	}
}

func (r *OrganzierRepo) GetOrganizerstats(ctx context.Context, organizerId int) (model.OrganizerStats, error) {
	sql := `SELECT COUNT(DISTINCT e.id) AS total_events, 
		COUNT(DISTINCT ue.user_id) AS total_attendees, 
		COALESCE(
			(COUNT(ue.user_id)::numeric / NULLIF(SUM(e.capacity) ,0)) * 100
		,0) AS avg_fill_rate
		FROM events e
		LEFT JOIN user_event ue ON ue.event_id = e.id
		WHERE e.organizer_id = $1`
	var stats model.OrganizerStats
	if err := r.db.QueryRow(ctx, sql, organizerId).Scan(&stats.TotalEvents, &stats.TotalAttandees, &stats.AvgFillRate); err != nil {
		return model.OrganizerStats{}, nil
	}
	return stats, nil

}
