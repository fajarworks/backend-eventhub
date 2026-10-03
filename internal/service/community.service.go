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
