package v1

import (
	"classconnect-api/internal/domain"
	"classconnect-api/internal/handler/http/v1/auth"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	services *domain.Services
	logger   *slog.Logger
}

func NewHandler(services *domain.Services, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitAPI(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		auth.NewHandler(h.services.Auth, h.logger).InitAPI(v1)
	}
}
