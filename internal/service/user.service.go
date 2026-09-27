package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/fajarworks/backend-eventhub/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserRepo(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetDetailUser(ctx context.Context, id int) (model.User, error) {
	return s.repo.FindUserById(ctx, id)
}

func (s *UserService) UpdateProfile(ctx context.Context, userId int, req dto.UpdateProfileRequest) (model.User, error) {
	var imagePath *string

	if req.PhotoProfile != nil {
		file := req.PhotoProfile
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			return model.User{}, apperror.ErrFileFormat
		}

		if file.Size > 2*1024*1024 {
			return model.User{}, apperror.ErrFileSize
		}
		filename := fmt.Sprintf("%d_%d%s", time.Now().UnixNano(), userId, path.Ext(req.PhotoProfile.Filename))
		filepath := path.Join("public", "images", filename)

		src, err := file.Open()
		if err != nil {
			return model.User{}, err
		}
		defer src.Close()
		data, err := io.ReadAll(src)
		if err != nil {
			return model.User{}, err
		}
		if err := os.WriteFile(filepath, data, 0644); err != nil {
			return model.User{}, err
		}

		var url string = "images" + filename
		imagePath = &url
	}
	dataUser, err := s.repo.UpdateProfile(ctx, userId, imagePath, req.Fullname, req.Location, req.JobPosition, req.Bio)
	if err != nil {
		return model.User{}, err
	}
	return dataUser, nil
}
