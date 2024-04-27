package auth

import (
	"context"
)

type Repository interface{}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) LogIn(ctx context.Context, dto LoginDTO) {

}

func (s *Service) SignUp(ctx context.Context, dto SignupDTO) {

}
