package v1

import (
	"classconnect-api/internal/domain"
	"classconnect-api/internal/handler/http/v1/auth"
	"classconnect-api/internal/handler/http/v1/group"
	"classconnect-api/internal/handler/http/v1/schedule"
	"classconnect-api/internal/handler/http/v1/subscriber"
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
		auth.NewHandler(h.services.Auth, h.logger).InitAPI(v1, h.services.Auth)
		group.NewHandler(h.services.Group, h.logger).InitAPI(v1, h.services.Auth)
		subscriber.NewHandler(h.services.Subscriber, h.logger).InitAPI(v1, h.services.Auth)
		schedule.NewHandler(h.services.Schedule, h.logger).InitAPI(v1, h.services.Auth)
	}
}
