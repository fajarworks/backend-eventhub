package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/middleware"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OrganizerRouter(r *gin.Engine, db *pgxpool.Pool) {
	organizerRouter := r.Group("/organizer")
	organizerRepo := repository.NewOrganizerRepo(db)
	organizerSerivce := service.NewOrganizerService(organizerRepo)
	organizerHandler := handler.NewOrganizerHandler(organizerSerivce)
	organizerRouter.GET("stats", middleware.AuthMiddleware, organizerHandler.GetOrganizerstats)

}
