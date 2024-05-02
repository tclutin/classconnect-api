package http

import (
	"classconnect-api/internal/config"
	"classconnect-api/internal/domain"
	"classconnect-api/internal/handler/http/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log/slog"
	"net/http"
)

func NewRouter(services *domain.Services, cfg *config.Config, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger(), gin.Recovery())

	if cfg.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	//TODO: one day add a swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	api := router.Group("/api")
	{
		v1.NewHandler(services, logger).InitAPI(api)
	}

	return router
}
