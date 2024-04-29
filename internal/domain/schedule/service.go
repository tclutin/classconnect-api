package schedule

import (
	"classconnect-api/internal/domain/group"
	"context"
)

type GroupRepository interface {
	UpdateGroup(ctx context.Context, group group.Group) error
}

type Repository interface {
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
