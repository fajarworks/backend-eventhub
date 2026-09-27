package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func EventRouter(r *gin.Engine, db *pgxpool.Pool) {
	eventRouter := r.Group("/events")
	eventRepo := repository.NewEventRepo(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	eventRouter.GET("", eventHandler.GetEvents)

}
