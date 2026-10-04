package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func (h *OrganizerHandler) CreateEvent(ctx *gin.Context) {
	var body dto.CreateEventRequest
	organizerId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "please login first",
		})
		return
	}
	if err := ctx.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), body.Title, path.Ext(body.Image.Filename))
	filepath := path.Join("public", "images", "events", filename)

	if err := ctx.SaveUploadedFile(&body.Image, filepath); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}

	var newSpeaker []dto.Speaker
	if body.NewSpeakers != "" {
		if err := json.Unmarshal([]byte(body.NewSpeakers), &newSpeaker); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: "invalid format new_speakers",
			})
			return
		}
	}
	eventId, err := h.service.CreateEvent(ctx, organizerId.(int), filepath, body, newSpeaker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: false,
		Data: gin.H{
			"eventId": eventId,
		},
		Message: "event created successfully",
	})

}
