package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	apperror "github.com/fajarworks/backend-eventhub/internal/errror"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
)

type CommunityHandler struct {
	service *service.CommunityService
}

func NewCommunityHandler(service *service.CommunityService) *CommunityHandler {
	return &CommunityHandler{
		service: service,
	}
}

func (h *CommunityHandler) GetDetailCommunity(ctx *gin.Context) {
	comId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	com, err := h.service.GetDetailCommunity(ctx.Request.Context(), comId)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: false,
		Data:    com,
		Message: "successfully retrieved community detail",
	})

}

func (h *CommunityHandler) GetCommunityMembers(ctx *gin.Context) {
	comId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{

			Success: false,
			Message: "terjadi kesalahan server",
		})
		return
	}
	data, err := h.service.GetCommunityMembers(ctx, comId)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, apperror.ErrNoRowsAffected) {
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "data not found",
			})

			return
		}
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    data,
		Message: "success retrieved community members",
	})
}
