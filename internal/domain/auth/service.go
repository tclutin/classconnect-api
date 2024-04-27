package auth

import (
	"context"
)

type Repository interface{}

type Service struct {
	repository Repository
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) LogIn(ctx context.Context, dto LoginDTO) (string, error) {
	return "", nil
}

func (s *Service) SignUp(ctx context.Context, dto SignupDTO) (string, error) {
	return "", nil
}

func (s *Service) GenerateToken(login string, password string) (string, error) {
	return "", nil
}
