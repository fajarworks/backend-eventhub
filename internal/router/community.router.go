package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/middleware"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func CommunityRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {

	comRouter := r.Group("/communities")
	comRepo := repository.NewCommunityRepo(db)
	comService := service.NewCommunityService(comRepo)
	comHandler := handler.NewCommunityHandler(comService)
	comRouter.GET("/:id", comHandler.GetDetailCommunity)
	comRouter.GET("/:id/members", comHandler.GetCommunityMembers)
	comRouter.POST("/:id/join", middleware.AuthMiddleware(rdb), comHandler.JoinCommunity)
	comRouter.DELETE("/:id/leave", middleware.AuthMiddleware(rdb), comHandler.LeaveCommunity)
	comRouter.GET("/popular", comHandler.GetPopularCommunity)
}
