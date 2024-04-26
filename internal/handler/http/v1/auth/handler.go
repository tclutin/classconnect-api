package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Service interface {
	LogIn(ctx context.Context)
	SignUp(ctx context.Context)
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

func (h *Handler) InitAPI(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", h.SignUp)
		authGroup.POST("/login", h.LogIn)
	}
}

func (h *Handler) LogIn(c *gin.Context) {

}

func (h *Handler) SignUp(c *gin.Context) {

}
