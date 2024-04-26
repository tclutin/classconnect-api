package v1

import (
	"classconnect-api/internal/handler/http/v1/auth"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) InitAPI(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		auth.NewHandler(nil, h.logger).InitAPI(v1)
	}
}
