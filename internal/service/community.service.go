package service

import (
	"context"

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
