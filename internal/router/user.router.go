package router

import (
	"path"

	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/middleware"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func UserRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	userRouter := r.Group("/user", middleware.AuthMiddleware)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserRepo(userRepo, rdb)
	userHandler := handler.NewUserHandler(userService)
	r.Static("images", path.Join("public", "images"))
	userRouter.GET("/detail", userHandler.GetDetailUser)
	userRouter.PATCH("/update", userHandler.UpdateUser)
	userRouter.POST("/change-password", userHandler.ChangePassword)
}
