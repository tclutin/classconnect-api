package schedule

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/internal/domain/group"
	"context"
	"errors"
	"strconv"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
}

type GroupRepository interface {
	GetGroupById(ctx context.Context, id string) (group.Group, error)
	UpdateGroup(ctx context.Context, group group.Group) error
}

type Repository interface {
	CreateSchedule(ctx context.Context, schedule UploadScheduleDTO, groupID uint64) error
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

	if err = s.ValidateSchedule(schedule); err != nil {
		return err
	}

	strGroupID := strconv.FormatUint(*user.GroupID, 10)

	group, err := s.groupRepo.GetGroupById(ctx, strGroupID)
	if err != nil {
		return err
	}

	if group.IsExistsSchedule {
		return ErrAlreadyExists
	}

	if err = s.repository.CreateSchedule(ctx, schedule, group.ID); err != nil {
		return err
	}

	group.IsExistsSchedule = true
	if err = s.groupRepo.UpdateGroup(ctx, group); err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateSchedule(schedule UploadScheduleDTO) error {
	if len(schedule.Weeks) != 1 && len(schedule.Weeks) != 2 {
		return ErrEvenGroup
	}

	if len(schedule.Weeks) == 1 {
		if len(schedule.Weeks[0].Days) > 6 {
			return ErrDaysCount
		}
	}

	if len(schedule.Weeks) == 2 {
		if len(schedule.Weeks[0].Days) > 6 || len(schedule.Weeks[1].Days) > 6 {
			return ErrDaysCount
		}

		if schedule.Weeks[0].IsEven && schedule.Weeks[1].IsEven {
			return ErrEvenGroup
		}

		if !schedule.Weeks[0].IsEven && !schedule.Weeks[1].IsEven {
			return ErrEvenGroup
		}
	}

	return nil
}
