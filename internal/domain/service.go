package domain

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/token"
)

type Services struct {
	Auth  *auth.Service
	Token *token.Service
}

func NewServices() *Services {
	return &Services{
		Auth:  auth.NewService(nil),
		Token: token.NewService(),
	}
}
