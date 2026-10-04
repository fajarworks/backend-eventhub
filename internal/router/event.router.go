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

func EventRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	eventRouter := r.Group("/events")
	eventRepo := repository.NewEventRepo(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	eventRouter.GET("", eventHandler.GetEvents)
	eventRouter.GET("/:id", eventHandler.GetEventDetail)
	eventRouter.POST("/:id/toggle-event", middleware.AuthMiddleware(rdb), eventHandler.ToggleJoinEvent)
	eventRouter.GET("/upcoming", eventHandler.UpcomingEvent)
	eventRouter.GET("/my-events", middleware.AuthMiddleware(rdb), eventHandler.GetEventsByUserId)
	eventRouter.POST("/:id/save", middleware.AuthMiddleware(rdb), eventHandler.SaveEvent)
	eventRouter.DELETE("/:id/remove", middleware.AuthMiddleware(rdb), eventHandler.RemoveSavedEvent)

}
