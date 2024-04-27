package domain

import (
	"classconnect-api/internal/domain/auth"
)

type Services struct {
	Auth *auth.Service
}

func NewServices() *Services {
	return &Services{
		Auth: auth.NewService(),
	}
}
