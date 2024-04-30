package schedule

import (
	"context"
)

type Repository interface {
	CreateSchedule(ctx context.Context, schedule UploadScheduleDTO) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) UploadSchedule(ctx context.Context, schedule UploadScheduleDTO, username string) error {

	return nil
}
