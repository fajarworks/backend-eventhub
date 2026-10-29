package service

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type CommunityService struct {
	repo *repository.CommunityRepo
}

func NewCommunityService(repo *repository.CommunityRepo) *CommunityService {
	return &CommunityService{
		repo: repo,
	}

}

func (s *CommunityService) GetDetailCommunity(ctx context.Context, comId int) (dto.CommunityResponse, error) {
	com, err := s.repo.GetDetailCommunity(ctx, comId)
	if err != nil {
		return dto.CommunityResponse{}, err
	}

	members, err := s.repo.CountCommunityMember(ctx, comId)
	if err != nil {
		return dto.CommunityResponse{}, err
	}
	categories, err := s.repo.GetCategoriesByCommunityId(ctx, comId)
	if err != nil {
		return dto.CommunityResponse{}, err
	}

	return dto.CommunityResponse{
		ID:          comId,
		Name:        com.Name,
		Description: com.Description,
		Categories:  categories,
		Members:     members,
	}, err

}

func (s *CommunityService) GetCommunityMembers(ctx context.Context, comId int) ([]dto.CommunityMembers, error) {
	data, err := s.repo.GetCommunityMembers(ctx, comId)
	members := make([]dto.CommunityMembers, 0, len(data))
	if err != nil {
		return nil, err
	}
	for _, member := range data {
		members = append(members, dto.CommunityMembers{
			ID:           member.ID,
			Name:         member.FullName,
			PhotoProfile: member.PhotoProfile,
			JobPosition:  member.JobPosition,
		})
	}
	return members, err
}

func (s *CommunityService) JoinCommunity(ctx context.Context, userId, comId int) error {
	err := s.repo.JoinCommunity(ctx, userId, comId)
	if err != nil {
		return err
	}
	return nil

}

func (s *CommunityService) LeaveCommunity(ctx context.Context, userId, comId int) error {
	err := s.repo.LeaveCommunity(ctx, userId, comId)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommunityService) GetPopularCommunity(ctx context.Context) ([]dto.PopularCommunities, error) {
	com, err := s.repo.GetPopularCommunity(ctx)
	if err != nil {
		return nil, err
	}
	var communities []dto.PopularCommunities
	for _, v := range com {
		communities = append(communities, dto.PopularCommunities{
			ID:             v.ID,
			Name:           v.Name,
			Image:          v.Image,
			Categories:     v.Categories,
			UpcomingEvents: v.UpcomingEvents,
			Members:        v.Members,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		})
	}
	return communities, nil
}

func (s *CommunityService) GetJoinedCommunities(ctx context.Context, userId int) ([]int, error) {
	return s.repo.GetJoinedCommunities(ctx, userId)
}

func (s *CommunityService) GetCommunitiesByUserId(ctx context.Context, userId int) ([]dto.UserCommunityResponse, error) {
	data, err := s.repo.GetCommunitiesByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	var communities []dto.UserCommunityResponse
	for _, com := range data {
		communities = append(communities, dto.UserCommunityResponse{
			ID:             com.ID,
			Image:          com.Image,
			Name:           com.Name,
			Categories:     com.Categories,
			Members:        com.Members,
			UpcomingEvents: com.UpcomingEvents,
		})
	}
	return communities, nil
}

func (s *CommunityService) GetAllCommunities(ctx context.Context, qp dto.CommunitiesFilter, endpoint string) ([]dto.CommunitiesResponse, dto.CommunityPageMetaData, error) {
	offset := (qp.Page - 1) * qp.Limit
	qp.Search = strings.TrimSpace(qp.Search)

	data, err := s.repo.GetAllCommunities(ctx, qp.Search, qp.CategoryID, qp.Limit, offset)
	if err != nil {
		log.Println(err.Error())
		return nil, dto.CommunityPageMetaData{}, err
	}
	var communities []dto.CommunitiesResponse
	for _, v := range data {
		communities = append(communities, dto.CommunitiesResponse{
			ID:             v.ID,
			Image:          v.Image,
			Name:           v.Name,
			Categories:     v.Categories,
			Members:        v.Members,
			UpcomingEvents: v.UpcomingEvents,
		})
	}
	totalData, err := s.repo.GetTotalData(ctx, qp.Search, qp.CategoryID)
	if err != nil {
		return nil, dto.CommunityPageMetaData{}, err
	}

	totalPage := (totalData + qp.Limit - 1) / qp.Limit

	var next, prev string
	categoryID := 0
	if qp.CategoryID != nil {
		categoryID = *qp.CategoryID
	}

	if qp.Page < totalPage {
		next = BuildPageURL(
			endpoint,
			qp.Search,
			categoryID,
			qp.Limit,
			qp.Page+1,
		)
	}

	if qp.Page > 1 && qp.Page <= totalPage {
		prev = BuildPageURL(
			endpoint,
			qp.Search,
			categoryID,
			qp.Limit,
			qp.Page-1,
		)
	}

	meta := dto.CommunityPageMetaData{
		Next:      next,
		Prev:      prev,
		TotalData: totalData,
	}
	return communities, meta, nil
}

func BuildPageURL(endPoint, search string, categoryID int, limit, page int) string {
	q := url.Values{}
	q.Set("page", strconv.Itoa(page))
	q.Set("limit", strconv.Itoa(limit))

	if search != "" {
		q.Set("search", search)
	}

	if categoryID != 0 {
		q.Set("category", strconv.Itoa(categoryID))
	}

	return fmt.Sprintf("%s?%s", endPoint, q.Encode())

}
