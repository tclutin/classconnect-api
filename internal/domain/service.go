package domain

import (
	"classconnect-api/internal/config"
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/repository/postgres"
)

type Services struct {
	Auth *auth.Service
}

func NewServices(config *config.Config, repositories *postgres.Repositories) *Services {
	return &Services{
		Auth: auth.NewService(config, repositories.User),
	}
}
