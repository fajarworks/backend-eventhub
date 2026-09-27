package router

import (
	"path"

	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/middleware"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserRouter(r *gin.Engine, db *pgxpool.Pool) {
	userRouter := r.Group("/user", middleware.AuthMiddleware)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserRepo(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r.Static("images", path.Join("public", "images"))
	userRouter.GET("/detail", userHandler.GetDetailUser)
	userRouter.PATCH("/update", userHandler.UpdateUser)
}
