package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestimonyRouter(r *gin.Engine, db *pgxpool.Pool) {
	testiRouter := r.Group("/testimonials")
	testiRepo := repository.NewTestimonyRepo(db)
	testiService := service.NewTestimonyService(testiRepo)
	testiHandler := handler.NewTestimonyHandler(testiService)

	testiRouter.GET("", testiHandler.GetTestimonials)

}
