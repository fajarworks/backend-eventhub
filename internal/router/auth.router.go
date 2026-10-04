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

func AuthRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	authRouter := r.Group("/auth")
	AuthRepo := repository.NewAuthRepo(db)
	AuthService := service.NewAuthService(AuthRepo, rdb)
	authhandler := handler.NewAuthHandler(AuthService)
	authRouter.POST("register", authhandler.Register)
	authRouter.POST("login", authhandler.Login)
	authRouter.POST("logout", middleware.AuthMiddleware(rdb), authhandler.Logout)
}
