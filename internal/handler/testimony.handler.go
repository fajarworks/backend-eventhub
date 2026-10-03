package handler

import (
	"log"
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
)

type TestimonyHandler struct {
	service *service.TestimonyService
}

func NewTestimonyHandler(service *service.TestimonyService) *TestimonyHandler {
	return &TestimonyHandler{
		service: service,
	}

}

// Testimoanials godoc
//
// @Summary			retrieves testimonials
// @description		get testimoanials for app
// @tags			Testimonials
// @produce			json
// @success			200 {object}	dto.Response
// @failure			500 {object}	dto.Response
// @router			/testimonials	[get]
func (h *TestimonyHandler) GetTestimonials(ctx *gin.Context) {
	testimonials, err := h.service.GetTestimonials(ctx)
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
		Data:    testimonials,
		Message: "success retrieved testimonials",
	})

}
