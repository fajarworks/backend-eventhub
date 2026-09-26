package router

import (
	"github.com/fajarworks/backend-eventhub/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool) {
	r.Use(middleware.Cors)

	AuthRouter(r, db)
}
