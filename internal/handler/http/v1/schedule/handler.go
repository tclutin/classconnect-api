package schedule

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/schedule"
	"classconnect-api/internal/handler/http/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	layerScheduleHandler = "handler.schedule."
)

type Service interface {
	UploadSchedule(ctx context.Context, schedule schedule.UploadScheduleDTO, username string) error
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

	fmt.Println(request.ToDTO())

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found in context"})
		return
	}

	//TODO: возможно username не лучшая идея... лучше передавать id группы
	err := h.service.UploadSchedule(context.Background(), request.ToDTO(), username.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}
