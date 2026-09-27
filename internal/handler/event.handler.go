package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(service *service.EventService) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

func (h *EventHandler) GetEvents(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "6"))
	location := ctx.Query("location")
	search := ctx.Query("search")
	sortBy := ctx.Query("sort")

	categoryId, _ := strconv.Atoi(ctx.Query("category"))

	events, err := h.service.GetEvents(ctx.Request.Context(), categoryId, location, search, sortBy, page, limit)
	if err != nil {
		log.Println("get events error:", err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "terjadi kesalahan server",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "berhasil mendapatkan data event",
		Data:    events,
	})
}
