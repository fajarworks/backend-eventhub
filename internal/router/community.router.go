package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CommunityRouter(r *gin.Engine, db *pgxpool.Pool) {

	comRouter := r.Group("/communites")
	comRepo := repository.NewCommunityRepo(db)
	comService := service.NewCommunityService(comRepo)
	comHandler := handler.NewCommunityHandler(comService)

	comRouter.GET("/:id", comHandler.GetDetailCommunity)

}
