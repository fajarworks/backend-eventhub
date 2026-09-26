package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/fajarworks/backend-eventhub/internal/dto"
	"github.com/fajarworks/backend-eventhub/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")

	if bearer == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "please login first",
		})
		return
	}

	result := strings.Split(bearer, " ")
	if len(result) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "invalid bearer token",
		})
		return
	}
	if result[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "invalid bearer token",
		})
		return
	}
	var token pkg.JWTClaims
	err := token.VerifyToken(result[1])
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Status:  false,
				Message: "invalid token",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Status:  false,
			Message: "terjadi kesalahan server",
		})
		return

	}
	ctx.Set("userId", token.Id)
	ctx.Next()
}
