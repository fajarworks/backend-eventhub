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
			Message: "server error occured",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Successfully retrieved events",
		Data:    events,
	})
}

func (h *EventHandler) ToggleJoinEvent(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid event id",
		})
		return
	}
	isJoined, err := h.service.ToggleJoinEvent(ctx.Request.Context(), userId.(int), eventId)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return

	}

	message := "user left event successfully"
	if isJoined {
		message = "user left event successfully"
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"is_Joined": isJoined,
		},
		Message: message,
	})

}
