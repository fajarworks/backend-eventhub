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
				Message: "user not found",
			})
		default:
			log.Println("get detail user error:", err)
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "server error occured",
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
		Message: "success retrieved user data",
	})

}

// UpdateUser
// @Summary Update User Information
// @Description Update information user
// @Tags User
// @Accept mpfd
// @Produce json
// @Security BearerAuth
// @Param photo_profile formData file false "Profile photo"
// @Param fullname formData string false "Full name"
// @Param job_position formData string false "Job position"
// @Param location formData string false "Location"
// @Param bio formData string false "Bio"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /user/update [patch]
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	var updateUserReq dto.UpdateProfileRequest

	if err := ctx.ShouldBindWith(&updateUserReq, binding.FormMultipart); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
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
				Message: "server error occured",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    user,
		Message: "user updated",
	})

}

func (h *UserHandler) ChangePassword(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
	}
	var body dto.ChangePassword
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	err := h.service.ChangePassword(ctx, userId.(int), body.OldPassword, body.NewPassword)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrWrongPass):
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "wrong old password",
			})
		case errors.Is(err, apperror.ErrEmptyField):
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "field can't be empty",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: false,
		Message: "success change password",
	})

}
