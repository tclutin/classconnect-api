package auth

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/handler/http/middleware"
	resp "classconnect-api/pkg/http"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	layerAuthHandler = "handler.auth."
)

type Service interface {
	LogIn(ctx context.Context, dto auth.LoginDTO) (string, error)
	SignUp(ctx context.Context, dto auth.SignupDTO) (string, error)
	GetUserByUsernameWithDetail(ctx context.Context, username string) (auth.UserDetailDTO, error)
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
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", h.SignUp)
		authGroup.POST("/login", h.LogIn)
		authGroup.GET("/me", middleware.AuthMiddleware(auth), h.Me)
	}
}

func (h *Handler) LogIn(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.NewAPIResponse(err))
		return
	}

	token, err := h.service.LogIn(c.Request.Context(), request.ToDTO())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewAPIResponse(err))
		return
	}

	c.JSON(http.StatusOK, TokenResponse{AccessToken: token})
}

// @Summary		Sign up
// @Description	Sign up with username, email, and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			input	body		SignupRequest	true	"User credentials"
// @Success		201		{object}	TokenResponse	"JWT token"
// @Failure		400		{object}	APIErrorResponse		"Invalid request payload"
// @Router			/auth/signup [post]
func (h *Handler) SignUp(c *gin.Context) {
	var request SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.NewAPIResponse(err))
		return
	}

	token, err := h.service.SignUp(c.Request.Context(), request.ToDTO())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewAPIResponse(err))
		return
	}

	c.JSON(http.StatusCreated, TokenResponse{AccessToken: token})
}

func (h *Handler) Me(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.NewAPIResponse("username not found in context"))
		return
	}

	userDetail, err := h.service.GetUserByUsernameWithDetail(c.Request.Context(), username.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewAPIResponse(err))
		return
	}

	c.JSON(http.StatusOK, ConvertUserDetailDTOToResponse(userDetail))
}
