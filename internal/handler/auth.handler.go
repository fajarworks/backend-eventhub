package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

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

// Login
// @Summary			Login to authentication
// @Description		login to access all features app
// @Tags			Auth
// @Accept       	json
// @Produce      	json
// @param			data	body	dto.LoginRequest	true "login woi"
// @success 		200 {object} dto.Response
// @failure			400 {object} dto.Response
// @failure			409 {object} dto.Response
// @failure			500 {object} dto.Response
// @router			/auth/login [post]
func (h *AuthHandler) Register(ctx *gin.Context) {
	var body dto.RegisterRequest

	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		log.Println("bind error:", err)

		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
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
			Message: "server error occured",
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
				Message: "server error occured",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    token,
		Message: "user login successfully",
	})

}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unathorized",
		})
		return
	}

	jti, exist := ctx.Get("jti")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "please login first",
		})
		return
	}
	exp, exist := ctx.Get("exp")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "please login first",
		})
		return
	}
	if err := h.service.Logout(ctx, userId.(int), jti.(string), exp.(time.Time)); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "logout success",
	})
}
