package auth

import (
	"classconnect-api/pkg/hash"
	"context"
	"time"
)

const (
	layerAuthService = "service.auth."
)

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByUsername(ctx context.Context, username string) (User, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) LogIn(ctx context.Context, dto LoginDTO) (string, error) {
	user, err := s.repository.GetUserByUsername(ctx, dto.Username)
	if err != nil {
		return "", ErrNotFound
	}

	if user.PasswordHash != hash.GenerateHash(dto.Password) {
		return "", ErrPasswordNotMatch
	}

	token, err := s.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) SignUp(ctx context.Context, dto SignupDTO) (string, error) {
	if _, err := s.repository.GetUserByUsername(ctx, dto.Username); err == nil {
		return "", ErrAlreadyExist
	}

	user := User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: hash.GenerateHash(dto.Password),
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
	return "пошёл нахуй", nil
}
