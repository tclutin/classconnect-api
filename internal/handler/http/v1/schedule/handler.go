package schedule

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/handler/http/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	layerScheduleHandler = "handler.schedule."
)

type Service interface {
	UploadSchedule(ctx context.Context, schedule UploadScheduleRequest, groupId uint64) error
}

type Handler struct {
	service Service
	logger  *slog.Logger
}

func NewHandler(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitAPI(router *gin.RouterGroup, auth *auth.Service) {
	scheduleGroup := router.Group("/schedules", middleware.AuthMiddleware(auth))
	{
		scheduleGroup.POST("/upload", h.UploadSchedule)
	}
}

func (h *Handler) UploadSchedule(c *gin.Context) {
	var request UploadScheduleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
