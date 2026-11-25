package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/fajarworks/backend-eventhub/internal/sse"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *service.EventService
	broker  *sse.Broker
}

func NewEventHandler(service *service.EventService, broker *sse.Broker) *EventHandler {
	return &EventHandler{
		service: service,
		broker:  broker,
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

func (h *EventHandler) StreamAttendees(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Data:    nil,
			Message: "invalid event id",
		})
		return
	}
	attendees, err := h.service.GetEventAttendees(ctx.Request.Context(), eventId)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "an error server occured",
		})
		return
	}
	w := ctx.Writer
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")

	ch := h.broker.Subscribe(eventId)
	defer h.broker.Unsubscribe(eventId, ch)

	send := func(count int) {
		fmt.Fprintf(w, "event: attendees\ndata: {\"count\":%d}\n\n", count)
		w.Flush()
	}
	send(attendees)
	ticker := time.NewTicker(25 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Request.Context().Done():
			return
		case count := <-ch:
			send(count)
		case <-ticker.C:
			fmt.Fprint(w, ": ping \n\n")
			w.Flush()
		}

	}

}

func (h *EventHandler) JoinEvent(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid event id",
		})
		return
	}

	if err := h.service.JoinEvent(ctx.Request.Context(), userId.(int), eventId); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user has successfully joined event",
	})
}

func (h *EventHandler) LeaveEvent(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid event id",
		})
		return
	}
	if err := h.service.LeaveEvent(ctx.Request.Context(), userId.(int), eventId); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user has successfully leave event",
	})
}

// Detail Event
// @Summary			retrieves an event detail
// @description		get a detail information of an event
// @tags			Events
// @produce			json
// @param			id path int	true "event ID"
// @success			200 {object}	dto.Response
// @failure			500 {object}	dto.Response
// @router			/events/{id}	[get]
func (h *EventHandler) GetEventDetail(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid event id",
		})
		return
	}
	event, err := h.service.GetDetail(ctx.Request.Context(), eventId)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    event,
		Message: "successfully retrieved event detail",
	})

}

func (h *EventHandler) UpcomingEvent(ctx *gin.Context) {
	events, err := h.service.GetUpcomingEvents(ctx)
	if err != nil {
		log.Println()
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    events,
		Message: "successfully get upcoming events",
	})

}

func (h *EventHandler) GetEventsByUserId(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}
	events, err := h.service.GetEventsByUserId(ctx, userId.(int))
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
		Data:    events,
		Message: "successfully retrieved events by user id",
	})

}

// SaveEvent godoc
//
// @Summary      saved an event
// @Description  Allows the authenticated user to save an event
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        id path int true "Event ID"
// @Security     BearerAuth
// @Router       /events/{id}/save [post]
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Failure      401 {object} dto.Response
func (h *EventHandler) SaveEvent(ctx *gin.Context) {
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
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "wrong event id",
		})
		return
	}

	if err := h.service.SaveEvent(ctx.Request.Context(), userId.(int), eventId); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user save event successfully",
	})
}

// RemoveSavedEvent godoc
//
// @Summary      remove saved a event
// @Description  Allows the authenticated user to remove a saved event
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        id path int true "Event ID"
// @Security     BearerAuth
// @Router       /events/{id}/remove [delete]
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Failure      401 {object} dto.Response
func (h *EventHandler) RemoveSavedEvent(ctx *gin.Context) {
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
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "wrong event id",
		})
		return
	}

	if err := h.service.RemoveSavedEvent(ctx, userId.(int), eventId); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user remove saved event successfully",
	})

}

func (h *EventHandler) GetJoinedEventIds(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}
	eventIds, err := h.service.GetJoinedEventIds(ctx.Request.Context(), userId.(int))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    eventIds,
		Message: "success get joined event Ids",
	})
}
func (h *EventHandler) GetSavedEventIds(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}
	eventIds, err := h.service.GetSavedEventIds(ctx.Request.Context(), userId.(int))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    eventIds,
		Message: "success get joined event Ids",
	})
}

func (h *EventHandler) GetEventsByCommunityId(ctx *gin.Context) {
	comId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Data:    nil,
			Message: "invalid id",
		})
		return
	}
	events, err := h.service.GetEventsByCommunityId(ctx, comId)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Message: "an error server occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    events,
		Message: "events retrieved successfully",
	})
}
