package auth

import (
	"classconnect-api/internal/config"
	"classconnect-api/pkg/hash"
	"context"
	"github.com/golang-jwt/jwt/v5"
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
	config     *config.Config
	repository Repository
}

func NewService(config *config.Config, repository Repository) *Service {
	return &Service{
		config:     config,
		repository: repository,
	}
}

func (s *Service) LogIn(ctx context.Context, dto LoginDTO) (string, error) {
	user, err := s.repository.GetUserByUsername(ctx, dto.Username)
	if err != nil {
		return "", ErrNotFound
	}

	if user.PasswordHash != hash.GenerateSha1Hash(dto.Password) {
		return "", ErrPasswordNotMatch
	}

	token, err := s.GenerateToken(user.Username, user.Email)
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
		PasswordHash: hash.GenerateSha1Hash(dto.Password),
		CreatedAt:    time.Now(),
	}

	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(user.Username, dto.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GenerateToken(username string, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username
	claims["email"] = email

	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
