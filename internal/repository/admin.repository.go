package repository

import (
	"context"
	"log"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepo struct {
	db *pgxpool.Pool
}

func NewAdminRepo(db *pgxpool.Pool) *AdminRepo {
	return &AdminRepo{
		db: db,
	}
}

func (r *AdminRepo) AdminOverview(ctx context.Context) (model.AdminOverview, error) {
	sql := `
		SELECT
			(SELECT COUNT(*) FROM users) AS total_users,
			(SELECT COUNT(*) FROM events) AS total_events,
			(SELECT COUNT(*) FROM communities) AS total_communities,
			COALESCE(
				ROUND(
					(
						SELECT COUNT(*)::numeric
						FROM user_event
					)
					/
					NULLIF(
						(SELECT SUM(capacity) FROM events),
						0
					) * 100
				),
				0
			) AS avg_fill_rate
	`

	var overview model.AdminOverview
	err := r.db.QueryRow(ctx, sql).Scan(&overview.TotalUser, &overview.TotalEvent, &overview.TotalCommunity, &overview.AvgFillRate)

	if err != nil {
		return model.AdminOverview{}, err
	}
	log.Println("data admin:", overview)
	return overview, err

}
