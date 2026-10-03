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

// JoinCommunity godoc
//
// @Summary      Join a community
// @Description  Allows the authenticated user to join a community
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        id path int true "Community ID"
// @Security     BearerAuth
// @Router       /communities/{id}/join [post]
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Failure      401 {object} dto.Response
func (h *CommunityHandler) JoinCommunity(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "please login first",
		})
		return
	}
	comId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "wrong community id",
		})
		return
	}
	if err := h.service.JoinCommunity(ctx.Request.Context(), userId.(int), comId); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "failed to join community",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user join event successfully",
	})

}
