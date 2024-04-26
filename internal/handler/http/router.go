package http

import (
	"classconnect-api/internal/config"
	"classconnect-api/internal/handler/http/v1"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func NewRouter(cfg *config.Config, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger(), gin.Recovery())

	if cfg.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	api := router.Group("/api")
	{
		v1.NewHandler(logger).InitAPI(api)
	}

	return router
}
