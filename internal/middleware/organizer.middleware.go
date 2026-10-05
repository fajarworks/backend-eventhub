package middleware

import (
	"net/http"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/gin-gonic/gin"
)

func OrganizerMiddleware(ctx *gin.Context) {
	role, exist := ctx.Get("role")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}
	if role.(string) != "organizer" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "you don't have permission to access this resource",
		})
		return
	}
	ctx.Next()
}
