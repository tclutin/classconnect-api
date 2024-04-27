package auth

import (
	"classconnect-api/internal/domain/auth"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const layer = "handler.auth."

type Service interface {
	LogIn(ctx context.Context, dto auth.LoginDTO) (string, error)
	SignUp(ctx context.Context, dto auth.SignupDTO) (string, error)
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
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info(layer+"LogIn", slog.String("username", request.Username))

	token, err := h.service.LogIn(context.Background(), request.ToDTO())
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, Token{AccessToken: token})
}

func (h *Handler) SignUp(c *gin.Context) {
	var request SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info(layer+"SignUp", slog.String("username", request.Username))

	token, err := h.service.SignUp(context.Background(), request.ToDTO())
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, Token{AccessToken: token})
}
