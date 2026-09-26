package handler

import (
	"log"
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) GetCategories(ctx *gin.Context) {
	categories, err := h.service.GetCategories(ctx.Request.Context())
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status:  false,
			Message: "terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Data:    categories,
		Message: "berhasil mendapatkan data category",
	})
}
