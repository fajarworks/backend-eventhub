package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
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
				Status:  false,
				Message: "user tidak ditemukan",
			})
		default:
			log.Println("get me error:", err)
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Status:  false,
				Message: "terjadi kesalahan server",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data: dto.UserResponse{
			Id:           user.Id,
			Role:         user.Role,
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
