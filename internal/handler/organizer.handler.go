package handler

import (
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
)

type OrganizerHandler struct {
	service *service.OrganizerService
}

func NewOrganizerHandler(service *service.OrganizerService) *OrganizerHandler {
	return &OrganizerHandler{
		service: service,
	}
}

// GetOrganizerstats godoc
//
// @Summary      Get organizer dashboard statistics
// @Description  Get statistic dashboard for organizer Authorized (total event, total attendee, and fill rate average).
// @Tags         Organizer
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response{data=dto.OrganizerStats}
// @Failure      401  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /organizer/stats [get]
func (h *OrganizerHandler) GetOrganizerstats(ctx *gin.Context) {
	organizerId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unathorized",
		})
		return
	}

	stats, err := h.service.GetOrganizerstats(ctx, organizerId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    stats,
		Message: "organizer stats retrieved successfully",
	})
}
