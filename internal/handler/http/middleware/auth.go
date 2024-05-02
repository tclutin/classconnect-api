package middleware

import (
	"classconnect-api/internal/domain/auth"
	resp "classconnect-api/pkg/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(auth *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.NewAPIErrorResponse("auth header is empty"))
			return
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.NewAPIErrorResponse("auth header is invalid"))
		}

		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp.NewAPIErrorResponse(err.Error()))
			return
		}

		c.Set("username", claims.Username)
	}
}
