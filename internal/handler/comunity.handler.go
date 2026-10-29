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
		Success: true,
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
			Message: "server error occured",
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

// JoinCommunity godoc
//
// @Summary      Leave a community
// @Description  Allows the authenticated user to leave a community
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        id path int true "Community ID"
// @Security     BearerAuth
// @Router       /communities/{id}/leave [delete]
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Failure      401 {object} dto.Response
func (h *CommunityHandler) LeaveCommunity(ctx *gin.Context) {
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
	if err := h.service.LeaveCommunity(ctx, userId.(int), comId); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user leave community successfully",
	})
}

// PopularCommunity  godoc
//
// @Description  Get Most Popular Communities
// @Tags         Community
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /communities/popular [get]
// @Success      200 {object} dto.Response
// @Failure      500 {object} dto.Response
func (h *CommunityHandler) GetPopularCommunity(ctx *gin.Context) {
	communities, err := h.service.GetPopularCommunity(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    communities,
		Message: "popular communities retrieved successfully",
	})
}

func (h *CommunityHandler) GetJoinedCommunities(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "please login first",
		})
		return
	}
	communityIds, err := h.service.GetJoinedCommunities(ctx.Request.Context(), userId.(int))
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
		Data:    communityIds,
		Message: "joined communities retrieved successfully",
	})
}

func (h *CommunityHandler) GetCommunitiesByUserId(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "please login first",
		})
		return
	}
	communities, err := h.service.GetCommunitiesByUserId(ctx.Request.Context(), userId.(int))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    communities,
		Message: "get communities by user id retrieved successfully",
	})
}

// GetAllCommunities godoc
// @Summary      Get all communities
// @Description  Retrieve a paginated list of communities with optional search and category filters
// @Tags         Communities
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"
// @Param        limit     query     int     false  "Items per page"
// @Param        category  query     int     false  "Filter by category ID"
// @Param        search    query     string  false  "Search community by name"
// @Success      200  {object}  dto.Response  "Communities retrieved successfully"
// @Failure      400  {object}  dto.Response  "Invalid query parameters"
// @Failure      500  {object}  dto.Response  "Internal server error"
// @Router       /communities [get]
func (h *CommunityHandler) GetAllCommunities(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "6"))
	if err != nil || limit < 1 {
		limit = 6
	}
	search := ctx.Query("search")
	var categoryID *int

	if category := ctx.Query("category"); category != "" {
		id, err := strconv.Atoi(category)
		if err != nil || id < 1 {
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "invalid category ID",
			})
			return
		}

		categoryID = &id
	}
	qp := dto.CommunitiesFilter{
		Search:     search,
		CategoryID: categoryID,
		Page:       page,
		Limit:      limit,
	}
	communites, meta, err := h.service.GetAllCommunities(ctx.Request.Context(), qp, ctx.Request.URL.Path)
	if err != nil {
		log.Println("err: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "an error server occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"data": communites,
			"meta": meta,
		},
		Message: "communities retrieved successfully",
	})
}
