package domain

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/repository/postgres"
)

type Services struct {
	Auth *auth.Service
}

func NewServices(repositories *postgres.Repositories) *Services {
	return &Services{
		Auth: auth.NewService(repositories.User),
	}
}
