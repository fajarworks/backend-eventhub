package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:8081"}
	if slices.Contains(allowedOrigins, ctx.GetHeader("Origin")) {
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
	}

	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Header("Access-Control-Allow-Method", "GET, OPTIONS, PATCH, PUT, POST")

	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}
