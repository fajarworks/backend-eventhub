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
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var body dto.RegisterRequest

	// Binding JSON
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		log.Println("bind error:", err)

		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "request body tidak valid",
		})
		return
	}

	if err := h.service.NewUser(ctx.Request.Context(), body); err != nil {
		log.Println("register error:", err)

		switch {
		case errors.Is(err, apperror.ErrEmptyField):
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: err.Error(),
			})

		case errors.Is(err, apperror.ErrInvalidEmailFormat):
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: err.Error(),
			})

		case errors.Is(err, apperror.ErrAlreadyExist):
			ctx.JSON(http.StatusConflict, dto.Response{
				Success: false,
				Message: err.Error(),
			})

		default:
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "user created",
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var body dto.LoginRequest

	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "terjadi kesalahan server",
		})
		return
	}

	token, err := h.service.User(ctx.Request.Context(), body)
	if err != nil {
		switch {

		case errors.Is(err, apperror.ErrEmptyField):
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: err.Error(),
			})
		case errors.Is(err, apperror.ErrWrongEmailPass):
			ctx.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: err.Error(),
			})
		default:
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    "",
				Message: "terjadi kesalahan server",
			})
		}
		return
	}
	ctx.JSON(http.StatusAccepted, dto.Response{
		Success: true,
		Data:    token,
		Message: "user berhasil login",
	})

}
