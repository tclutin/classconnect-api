package domain

import (
	"classconnect-api/internal/config"
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/group"
	"classconnect-api/internal/repository/postgres"
)

type Services struct {
	Auth  *auth.Service
	Group *group.Service
}

func NewServices(config *config.Config, repositories *postgres.Repositories) *Services {
	authService := auth.NewService(config, repositories.User)
	groupService := group.NewService(repositories.Group, repositories.User)
	return &Services{
		Auth:  authService,
		Group: groupService,
	}
}
