package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/middleware"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AdminRouter(r *gin.Engine, db *pgxpool.Pool) {
	adminRouter := r.Group("/admin")
	adminRepo := repository.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)
	adminRouter.GET("overview", middleware.AuthMiddleware, middleware.AdminMiddleware, adminHandler.GetAdminOverviews)

}
