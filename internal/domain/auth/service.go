package auth

import (
	"context"
	"time"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByUsername(ctx context.Context, username string) (User, error)
}

type Service struct {
	repository Repository
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) LogIn(ctx context.Context, dto LoginDTO) (string, error) {
	if _, err := s.repository.GetUserByUsername(ctx, dto.Username); err != nil {
		return "", ErrNotFound
	}

	token, err := s.GenerateToken(dto.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) SignUp(ctx context.Context, dto SignupDTO) (string, error) {
	if _, err := s.repository.GetUserByUsername(ctx, dto.Username); err != nil {
		return "", ErrAlreadyExist
	}

	user := User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: dto.Password,
		CreatedAt:    time.Now(),
	}

	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GenerateToken(username string) (string, error) {
	return "", nil
}
