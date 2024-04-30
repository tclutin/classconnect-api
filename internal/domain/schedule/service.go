package schedule

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/group"
	"context"
	"errors"
)

type GroupRepository interface {
	GetGroupById(ctx context.Context, id string) (group.Group, error)
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
}

type Repository interface {
	CreateSchedule(ctx context.Context, schedule UploadScheduleDTO, groupID *uint64) error
}

type Service struct {
	repository Repository
	userRepo   UserRepository
	groupRepo  GroupRepository
}

func NewService(repository Repository, userRepository UserRepository, groupRepository group.Repository) *Service {
	return &Service{
		repository: repository,
		userRepo:   userRepository,
		groupRepo:  groupRepository,
	}
}

func (s *Service) UploadSchedule(ctx context.Context, schedule UploadScheduleDTO, username string) error {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return auth.ErrNotFound
	}

	if user.GroupID == nil {
		return errors.New("user does not have a group")
	}

	if len(schedule.Weeks) != 1 && len(schedule.Weeks) != 2 {
		return ErrEvenGroup
	}

	if err = s.repository.CreateSchedule(ctx, schedule, user.GroupID); err != nil {
		return err
	}

	return nil
}
