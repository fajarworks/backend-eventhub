package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fajarworks/backend-eventhub/internal/dto"
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
