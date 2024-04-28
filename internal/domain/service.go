package domain

import (
	"classconnect-api/internal/config"
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/group"
	"classconnect-api/internal/domain/subscriber"
	"classconnect-api/internal/repository/postgres"
)

type Services struct {
	Auth       *auth.Service
	Group      *group.Service
	Subscriber *subscriber.Service
}

func NewServices(config *config.Config, repositories *postgres.Repositories) *Services {
	authService := auth.NewService(config, repositories.User)
	groupService := group.NewService(repositories.Group, repositories.User)
	subscriberService := subscriber.NewService(repositories.Subscriber)
	return &Services{
		Auth:       authService,
		Group:      groupService,
		Subscriber: subscriberService,
	}
}
