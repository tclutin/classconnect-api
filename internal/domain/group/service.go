package group

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/pkg/hash"
	"context"
	"errors"
	"time"
)

const (
	layerGroupService = "service.group."
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
	UpdateUser(ctx context.Context, user auth.User) error
}

type Repository interface {
	CreateGroup(ctx context.Context, group Group) error
	GetGroupByName(ctx context.Context, name string) (Group, error)
	GetAllGroups(ctx context.Context) ([]Group, error)
}

type Service struct {
	repository     Repository
	userRepository UserRepository
}

func NewService(repository Repository, userRepository UserRepository) *Service {
	return &Service{
		repository:     repository,
		userRepository: userRepository,
	}
}

func (s *Service) CreateGroup(ctx context.Context, username string, name string) (Group, error) {
	if _, err := s.repository.GetGroupByName(ctx, name); err == nil {
		return Group{}, ErrAlreadyExists
	}

	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return Group{}, auth.ErrNotFound
	}

	if user.GroupID != nil {
		return Group{}, errors.New("you already have group")
	}

	createGroup := Group{
		Name:      name,
		Code:      s.GenerateName(4),
		CreatedAt: time.Now(),
	}

	if err = s.repository.CreateGroup(ctx, createGroup); err != nil {
		return Group{}, err
	}

	group, err := s.repository.GetGroupByName(ctx, name)
	if err != nil {
		return Group{}, err
	}

	user.GroupID = &group.ID

	if err = s.userRepository.UpdateUser(ctx, user); err != nil {
		return Group{}, err
	}

	return group, nil
}

func (s *Service) GetAllGroups(ctx context.Context) ([]Group, error) {
	groups, err := s.repository.GetAllGroups(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *Service) GenerateName(size int64) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	alias := make([]rune, size)
	for i := range alias {
		alias[i] = chars[hash.NewCryptoRand(int64(len(chars)))]
	}
	return string(alias)
}
