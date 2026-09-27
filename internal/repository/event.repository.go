package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepo struct {
	db *pgxpool.Pool
}

type EventFilter struct {
	CategoryId int
	Location   string
	Search     string
	SortBy     string
	Limit      int
	Offset     int
}

func NewEventRepo(db *pgxpool.Pool) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (r *EventRepo) GetEvents(ctx context.Context, filter EventFilter) ([]model.Event, error) {
	baseQuery := `
		SELECT e.id, e.title, e.image, e.location, e.capacity, e.start_time
		FROM events e
	`

	var joins []string
	var conditions []string
	var args []any
	argPos := 1

	if filter.SortBy == "popular" {
		joins = append(joins, "LEFT JOIN user_event ue ON ue.event_id = e.id")
	}

	for _, j := range joins {
		baseQuery += j
	}

	if filter.CategoryId != 0 {
		conditions = append(conditions, fmt.Sprintf(`
			EXISTS (
				SELECT 1 FROM event_category ec
				WHERE ec.event_id = e.id AND ec.category_id = $%d
			)
		`, argPos))
		args = append(args, filter.CategoryId)
		argPos++
	}

	if filter.Location != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(e.location) LIKE LOWER($%d)", argPos))
		args = append(args, "%"+filter.Location+"%")
		argPos++
	}

	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(e.title) LIKE LOWER($%d)", argPos))
		args = append(args, "%"+filter.Search+"%")
		argPos++
	}

	if filter.SortBy == "upcoming" {
		conditions = append(conditions, "e.start_time >= now()")
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	switch filter.SortBy {
	case "popular":
		baseQuery += " GROUP BY e.id ORDER BY COUNT(DISTINCT ue.user_id) DESC"
	case "upcoming":
		baseQuery += " ORDER BY e.start_time ASC"
	default:
		baseQuery += " ORDER BY e.id"
	}

	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, filter.Limit, filter.Offset)
	log.Println(baseQuery)
	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Image, &e.Location, &e.Capacity, &e.StartTime); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepo) GetCategoriesByEventId(ctx context.Context, eventId int) ([]string, error) {
	sql := `
		SELECT c.name FROM categories c
		JOIN event_category ec ON ec.category_id = c.id
		WHERE ec.event_id = $1
	`

	rows, err := r.db.Query(ctx, sql, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		categories = append(categories, name)
	}
	return categories, rows.Err()
}

func (r *EventRepo) GetEventAttendees(ctx context.Context, eventId int) (int, error) {
	sql := "SELECT COUNT(*) FROM user_event WHERE event_id = $1"
	var attendees int
	if err := r.db.QueryRow(ctx, sql, eventId).Scan(&attendees); err != nil {
		return 0, err
	}
	return attendees, nil

}

func (r *EventRepo) JoinEvent(ctx context.Context, userId, eventId int) error {
	sql := "INSERT INTO user_event (user_id, event_id) VALUES ($1, $2)"
	args := []any{userId, eventId}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil

}
