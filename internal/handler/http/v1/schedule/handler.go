package schedule

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/schedule"
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
	UploadSchedule(ctx context.Context, schedule schedule.UploadScheduleDTO, username string) error
	GetScheduleForDay(ctx context.Context, subID uint64) ([]schedule.SubjectDTO, error)
	DeleteSchedule(ctx context.Context, username string) error
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
		scheduleGroup.GET("", h.GetScheduleForDay)
		scheduleGroup.DELETE("", h.DeleteSchedule)
	}
}

func (h *Handler) UploadSchedule(c *gin.Context) {
	var request UploadScheduleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found in context"})
		return
	}

	err := h.service.UploadSchedule(context.Background(), request.ToDTO(), username.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}

func (h *Handler) GetScheduleForDay(c *gin.Context) {
	var request GetScheduleForDayRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule, err := h.service.GetScheduleForDay(c.Request.Context(), request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func (h *Handler) DeleteSchedule(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found in context"})
		return
	}

	if err := h.service.DeleteSchedule(c.Request.Context(), username.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}
