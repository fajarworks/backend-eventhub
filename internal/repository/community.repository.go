package repository

import (
	"context"
	"fmt"
	"strings"

	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CommunityRepo struct {
	db *pgxpool.Pool
}

func NewCommunityRepo(db *pgxpool.Pool) *CommunityRepo {
	return &CommunityRepo{
		db: db,
	}

}

func (r *CommunityRepo) GetDetailCommunity(ctx context.Context, comId int) (model.Community, error) {
	sql := `SELECT id, name, description FROM communities WHERE id = $1`
	var community model.Community
	err := r.db.QueryRow(ctx, sql, comId).Scan(&community.ID, &community.Name, &community.Description)
	if err != nil {
		return model.Community{}, err
	}
	return community, nil

}

func (r *CommunityRepo) CountCommunityMember(ctx context.Context, comId int) (int, error) {
	sql := `SELECT COUNT(*) FROM user_community WHERE community_id = $1`
	var members int
	if err := r.db.QueryRow(ctx, sql, comId).Scan(&members); err != nil {
		return 0, err
	}
	return members, nil
}

func (r *CommunityRepo) GetCategoriesByCommunityId(ctx context.Context, comId int) ([]string, error) {
	sql := `
		SELECT c.name FROM categories c
		JOIN community_category cc ON cc.category_id = c.id
		WHERE cc.community_id = $1`

	rows, err := r.db.Query(ctx, sql, comId)
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

func (r *CommunityRepo) GetCommunityMembers(ctx context.Context, comId int) ([]model.CommunityMember, error) {
	sql := `SELECT u.id, u.fullname, u.job_position, COALESCE( u.photo_profile,'')
	FROM user_community uc
	JOIN users u ON uc.user_id = u.id
	WHERE uc.community_id = $1`

	rows, err := r.db.Query(ctx, sql, comId)
	if err != nil {
		return nil, err

	}
	defer rows.Close()

	var members []model.CommunityMember

	for rows.Next() {
		var member model.CommunityMember
		if err := rows.Scan(&member.ID, &member.FullName, &member.JobPosition, &member.PhotoProfile); err != nil {
			return nil, err
		}
		members = append(members, member)

	}
	return members, nil
}

func (r *CommunityRepo) JoinCommunity(ctx context.Context, userId, comId int) error {
	sql := `INSERT INTO user_community (user_id, community_id) VALUES($1, $2)`
	cmdTag, err := r.db.Exec(ctx, sql, userId, comId)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return apperror.ErrNoRowsAffected
	}
	return nil
}

func (r *CommunityRepo) LeaveCommunity(ctx context.Context, userId, comId int) error {
	sql := `DELETE FROM user_community WHERE user_id = $1 AND community_id = $2`
	cmdTag, err := r.db.Exec(ctx, sql, userId, comId)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return apperror.ErrNoRowsAffected
	}
	return nil
}

func (r *CommunityRepo) GetPopularCommunity(ctx context.Context) ([]model.PopularCommunities, error) {
	sql := `SELECT
		communities.id,
		communities.name,
		communities.image,
		ARRAY_AGG(DISTINCT categories.name) AS categories,
		COUNT(DISTINCT user_community.user_id) AS members_count,
		COUNT(DISTINCT events.id) AS upcoming_events_count
	FROM communities
	LEFT JOIN community_category ON communities.id = community_category.community_id
	LEFT JOIN categories ON categories.id = community_category.category_id
	LEFT JOIN user_community ON communities.id = user_community.community_id
	LEFT JOIN events ON events.community_id = communities.id AND events.start_time >= now()
	GROUP BY communities.id
	ORDER BY members_count DESC
	LIMIT 4`
	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var communities []model.PopularCommunities
	for rows.Next() {
		var community model.PopularCommunities

		err := rows.Scan(&community.ID, &community.Name, &community.Image, &community.Categories, &community.Members, &community.UpcomingEvents)
		if err != nil {
			return nil, err

		}
		communities = append(communities, community)
	}
	return communities, nil
}

func (r *CommunityRepo) GetJoinedCommunities(ctx context.Context, userId int) ([]int, error) {
	sql := `SELECT community_id FROM user_community WHERE user_id = $1`

	rows, err := r.db.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}
	communityIds := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		communityIds = append(communityIds, id)
	}
	return communityIds, err

}

func (r *CommunityRepo) GetCommunitiesByUserId(ctx context.Context, userId int) ([]model.CommunityResponse, error) {
	sql := `SELECT
			c.id, c.name, c.image,
			COALESCE((
				SELECT ARRAY_AGG(cat.name ORDER BY cat.name)
				FROM community_category cc
				JOIN categories cat ON cat.id = cc.category_id
				WHERE cc.community_id = c.id
			), '{}') AS categories,
			(SELECT COUNT(*) FROM user_community m WHERE m.community_id = c.id) AS members_count,
			(SELECT COUNT(*) FROM events e
				WHERE e.community_id = c.id AND e.start_time >= now()) AS upcoming_events_count
		FROM user_community uc
		JOIN communities c ON c.id = uc.community_id
		WHERE uc.user_id = $1
		ORDER BY c.name`
	rows, err := r.db.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}
	data := make([]model.CommunityResponse, 0)
	for rows.Next() {
		var com model.CommunityResponse
		if err := rows.Scan(&com.ID, &com.Name, &com.Image, &com.Categories, &com.Members, &com.UpcomingEvents); err != nil {
			return nil, err
		}
		data = append(data, com)
	}
	return data, nil
}

func (r *CommunityRepo) GetAllCommunities(ctx context.Context, search string, categoryID *int, limit, offset int) ([]model.Community, error) {
	var query strings.Builder
	var args []any

	query.WriteString(
		`	SELECT 
    com.id,
    com.name,
		com.image,
    (
        SELECT COUNT(*)
        FROM user_community uc
        WHERE uc.community_id = com.id
    ) AS members,
    (
        SELECT array_agg(c.name)
        FROM community_category cc
        JOIN categories c ON c.id = cc.category_id
        WHERE cc.community_id = com.id
    ) AS categories,
    (
        SELECT COUNT(*)
        FROM events e
        WHERE e.community_id = com.id AND e.start_time >= now()
    ) AS upcoming_events
FROM communities com WHERE 1=1`)

	if categoryID != nil {
		args = append(args, *categoryID)
		query.WriteString(fmt.Sprintf(`
					AND EXISTS ( SELECT 1 FROM community_category cc
						WHERE cc.community_id = com.id
							AND cc.category_id = $%d)`, len(args)))
	}

	if search != "" {
		args = append(args, "%"+search+"%")
		query.WriteString(fmt.Sprintf(`
				AND LOWER(com.name) LIKE LOWER($%d)`, len(args)))
	}

	args = append(args, limit)
	query.WriteString(fmt.Sprintf(` LIMIT $%d`, len(args)))

	args = append(args, offset)
	query.WriteString(fmt.Sprintf(` OFFSET $%d`, len(args)))

	rows, err := r.db.Query(ctx, query.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var communities []model.Community

	for rows.Next() {
		var com model.Community
		if err := rows.Scan(&com.ID, &com.Name, &com.Image, &com.Members, &com.Categories, &com.UpcomingEvents); err != nil {
			return nil, err

		}
		communities = append(communities, com)
	}
	return communities, nil
}

func (r *CommunityRepo) GetTotalData(ctx context.Context, search string, CategoryID *int) (int, error) {
	var q strings.Builder
	var args []any

	q.WriteString(`SELECT COUNT(*) FROM communities c  WHERE 1=1`)
	if CategoryID != nil {
		args = append(args, *CategoryID)
		q.WriteString(fmt.Sprintf(`
		AND EXISTS (
		SELECT 1 FROM community_category cc
		WHERE cc.community_id = c.id AND
		cc.category_id = $%d
		)`, len(args)))

	}
	if search != "" {
		args = append(args, "%"+search+"%")
		q.WriteString(fmt.Sprintf(`
		AND LOWER(c.name) LIKE LOWER($%d)`, len(args)))
	}
	var totalData int
	if err := r.db.QueryRow(ctx, q.String(), args...).Scan(&totalData); err != nil {
		return 0, err
	}
	return totalData, nil
}
