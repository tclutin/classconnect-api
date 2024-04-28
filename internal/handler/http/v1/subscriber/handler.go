package subscriber

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/subscriber"
	"classconnect-api/internal/handler/http/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

const (
	layerSubscriberHandler = "handler.subscriber."
)

type Service interface {
	CreateDeviceSubscriber(ctx context.Context, deviceId uint64) error
	CreateTelegramSubscriber(ctx context.Context, deviceId uint64) error
	EnableNotificationSubscriber(ctx context.Context, subId uint64, isNotification bool) error
	GetSubscriberByDeviceId(ctx context.Context, deviceId uint64) (subscriber.Subscriber, error)
	GetSubscriberByChatId(ctx context.Context, chatId uint64) (subscriber.Subscriber, error)
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
	subscriberGroup := router.Group("/subscribers", middleware.AuthMiddleware(auth))
	{
		subscriberGroup.POST("/device", h.CreateDeviceSubscriber)
		subscriberGroup.POST("/telegram", h.CreateTelegramSubscriber)
		subscriberGroup.GET("/device/:deviceId", h.GetSubscriberByDeviceId)
		subscriberGroup.GET("/telegram/:chatId", h.GetSubscriberChatId)
		subscriberGroup.PATCH("/:subscriberId", h.EnableNotification)
	}
}

func (h *Handler) CreateTelegramSubscriber(c *gin.Context) {
	var request CreateTelegramSubscriberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateTelegramSubscriber(context.Background(), request.ChatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(request.ChatId)

	c.JSON(http.StatusCreated, gin.H{"status": "Successfully"})
}

func (h *Handler) CreateDeviceSubscriber(c *gin.Context) {
	var request CreateDeviceSubscriberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateDeviceSubscriber(context.Background(), request.DeviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Successfully"})
}

func (h *Handler) EnableNotification(c *gin.Context) {
	var request EnableNotificationSubscriberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("subscriberId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.EnableNotificationSubscriber(context.Background(), uint64(id), request.Notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}

func (h *Handler) GetSubscriberByDeviceId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("deviceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.service.GetSubscriberByDeviceId(context.Background(), uint64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sub)
}

func (h *Handler) GetSubscriberChatId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("chatId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.service.GetSubscriberByChatId(context.Background(), uint64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sub)
}
