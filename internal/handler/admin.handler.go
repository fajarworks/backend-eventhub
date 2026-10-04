package handler

import (
	"log"
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *service.AdminService
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

// GetAdminOverviews godoc
// @Summary      Get admin dashboard overview
// @Description  Retrieves overview data for the admin dashboard (platform-wide summary statistics). Requires authentication and admin role.
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response{data=dto.AdminOverview}
// @Failure      401  {object}  dto.Response  "Unauthorized - token is missing, invalid, or expired"
// @Failure      403  {object}  dto.Response  "Forbidden - user is not an admin"
// @Failure      500  {object}  dto.Response  "Internal server error"
// @Router       /admin/overview [get]
func (h *AdminHandler) GetAdminOverviews(ctx *gin.Context) {
	overview, err := h.service.AdminOverview(ctx)
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
		Data:    overview,
		Message: "data admin overviews retrieved successfully",
	})

}
