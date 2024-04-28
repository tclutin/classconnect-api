package group

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/group"
	"classconnect-api/internal/handler/http/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	layerGroupHandler = "handler.group."
)

type Service interface {
	CreateGroup(ctx context.Context, username string, name string) (group.Group, error)
	GetAllGroups(ctx context.Context) ([]group.Group, error)
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
	groupGroup := router.Group("/groups", middleware.AuthMiddleware(auth))
	{
		groupGroup.POST("", h.CreateGroup)
		//groupGroup.POST("/:groupID/join", h.JoinToGroup)
		groupGroup.GET("", h.GetAllGroups)
		//groupGroup.DELETE("/:groupID", middleware.AuthMiddleware(auth), h.DeleteGroup)
	}
}

func (h *Handler) CreateGroup(c *gin.Context) {
	var request CreateGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found in context"})
		return
	}

	createdGroup, err := h.service.CreateGroup(context.Background(), username.(string), request.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info(layerGroupHandler+"CreateGroup", slog.String("name", request.Name))

	c.JSON(http.StatusCreated, createdGroup)
}

func (h *Handler) GetAllGroups(c *gin.Context) {
	groups, err := h.service.GetAllGroups(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info(layerGroupHandler + "GetAllGroups")

	c.JSON(http.StatusOK, groups)
}

func (h *Handler) DeleteGroup(c *gin.Context) {

}

func (h *Handler) JoinToGroup(c *gin.Context) {

}
