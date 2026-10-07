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

// CreateEvent godoc
// @Summary      Create a new event
// @Description  Creates a new event with an uploaded image, and optionally assigns existing categories, existing speakers, and new speakers (as a JSON array string). Requires authentication as an organizer.
// @Tags         Organizer
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        title          formData  string  true   "Event title"
// @Param        description    formData  string  false  "Event description"
// @Param        location       formData  string  true   "Event location"
// @Param        capacity       formData  int     true   "Maximum number of attendees"
// @Param        start_time     formData  string  true   "Event start time"
// @Param        end_time       formData  string  true   "Event end time"
// @Param        community_id   formData  int     false  "Community ID this event belongs to"
// @Param        categories     formData  []int   false  "Existing category IDs to assign" collectionFormat(multi)
// @Param        speakers       formData  []int   false  "Existing speaker IDs to assign" collectionFormat(multi)
// @Param        new_speakers   formData  string  false  "[{\"name\":\"John Doe\",\"job_position\":\"CTO\"}]"
// @Param        image          formData  file    true   "Event cover image"
// @Success      200  {object}  dto.Response{data=object{eventId=int}}  "Event created successfully"
// @Failure      400  {object}  dto.Response  "Invalid request body or new_speakers format"
// @Failure      401  {object}  dto.Response  "Unauthorized - please login first"
// @Failure      500  {object}  dto.Response  "Internal server error"
// @Router       /organizer/create-event [post]
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
