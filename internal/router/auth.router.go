package router

import (
	"github.com/fajarworks/backend-eventhub/internal/handler"
	"github.com/fajarworks/backend-eventhub/internal/repository"
	"github.com/fajarworks/backend-eventhub/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRouter(r *gin.Engine, db *pgxpool.Pool) {
	authRouter := r.Group("/auth")
	AuthRepo := repository.NewAuthRepo(db)
	AuthService := service.NewAuthService(AuthRepo)
	authhandler := handler.NewAuthHandler(AuthService)
	authRouter.POST("register", authhandler.Register)
	authRouter.POST("login", authhandler.Login)
}
