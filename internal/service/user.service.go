package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/model"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	repo *repository.UserRepo
	rdb  *redis.Client
}

func NewUserRepo(repo *repository.UserRepo, rdb *redis.Client) *UserService {
	return &UserService{
		repo: repo,
		rdb:  rdb,
	}
}

func (s *UserService) GetDetailUser(ctx context.Context, id int) (model.User, error) {
	key := fmt.Sprintf("eventhub:user%d", id)
	if res, err := s.rdb.Get(ctx, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			log.Println("key doesn't exist")
		} else {
			log.Println(err.Error())

		}
	} else {
		var data model.User
		if err := json.Unmarshal([]byte(res), &data); err != nil {
			log.Println("failed to parse: ", err.Error())
		} else {
			return data, nil
		}
	}

	person, err := s.repo.FindUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	if res, err := json.Marshal(person); err != nil {
		log.Println("failed to stringfy :", err.Error())

	} else {
		if err := s.rdb.Set(ctx, key, string(res), 0).Err(); err != nil {
			log.Println("failed set to redis", err.Error())
		}
	}
	return person, nil

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
