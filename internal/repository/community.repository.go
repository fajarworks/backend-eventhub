package repository

import (
	"context"

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
