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
	GetGroupById(ctx context.Context, groupID string) (group.Group, error)
	JoinToGroup(ctx context.Context, groupId string, subId uint64, code string) error
	LeaveFromGroup(ctx context.Context, groupId string, subId uint64) error
	DeleteGroup(ctx context.Context, groupId string) error
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
		groupGroup.GET("", h.GetAllGroups)
		groupGroup.DELETE("", h.DeleteGroup)
		groupGroup.GET("/:groupID", h.GetGroupById)
		groupGroup.POST("/:groupID/join", h.JoinToGroup)
		groupGroup.POST("/:groupID/leave", h.LeaveFromGroup)
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

	c.JSON(http.StatusOK, ConvertGroupsToResponse(groups))
}

func (h *Handler) JoinToGroup(c *gin.Context) {
	var request JoinToGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupId := c.Param("groupID")

	err := h.service.JoinToGroup(context.Background(), groupId, request.ID, request.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}

func (h *Handler) LeaveFromGroup(c *gin.Context) {
	var request LeaveFromGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupId := c.Param("groupID")

	if err := h.service.LeaveFromGroup(c.Request.Context(), groupId, request.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}

func (h *Handler) DeleteGroup(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username not found in context"})
		return
	}

	if err := h.service.DeleteGroup(c.Request.Context(), username.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully"})
}

func (h *Handler) GetGroupById(c *gin.Context) {
	groupId := c.Param("groupID")

	group, err := h.service.GetGroupById(c.Request.Context(), groupId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ConvertGroupToResponse(group))
}
