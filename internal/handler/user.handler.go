package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}

}

func (h *UserHandler) GetDetailUser(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	user, err := h.service.GetDetailUser(ctx.Request.Context(), userId.(int))
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			ctx.JSON(http.StatusNotFound, dto.Response{
				Success: false,
				Message: "user tidak ditemukan",
			})
		default:
			log.Println("get detail user error:", err)
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "terjadi kesalahan server",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: dto.UserResponse{
			Id:           user.Id,
			Role:         user.Role,
			Fullname:     user.Fullname,
			Email:        user.Email,
			PhotoProfile: user.PhotoProfile,
			JobPosition:  user.JobPosition,
			Location:     user.Location,
			Bio:          user.Bio,
			CreatedAt:    user.CreatedAt,
		},
		Message: "berhasil mendapatkan data user",
	})

}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	var updateUserReq dto.UpdateProfileRequest

	if err := ctx.ShouldBindWith(&updateUserReq, binding.FormMultipart); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "terjadi kesalahan sistem",
		})
		return
	}

	user, err := h.service.UpdateProfile(ctx.Request.Context(), userId.(int), updateUserReq)
	if err != nil {
		log.Println(err.Error())
		switch {
		case errors.Is(err, apperror.ErrFileFormat):
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: err.Error(),
			})
		case errors.Is(err, apperror.ErrFileSize):
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: err.Error(),
			})
		default:

			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "terjadi kesalahan server",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: dto.UserResponse{
			Id:           user.Id,
			PhotoProfile: user.PhotoProfile,
			Fullname:     user.Fullname,
			JobPosition:  user.JobPosition,
			Location:     user.Location,
			Bio:          user.Bio,
			UpdatedAt:    user.UpdatedAt,
		},
		Message: "berhasil update user",
	})

}
