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

func OrganizerRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	organizerRouter := r.Group("/organizer", middleware.AuthMiddleware(rdb), middleware.OrganizerMiddleware)
	organizerRepo := repository.NewOrganizerRepo()
	organizerSerivce := service.NewOrganizerService(organizerRepo, db)
	organizerHandler := handler.NewOrganizerHandler(organizerSerivce)
	organizerRouter.GET("stats", organizerHandler.GetOrganizerstats)
	organizerRouter.POST("/create-event", organizerHandler.CreateEvent)

}
