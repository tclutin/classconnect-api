package postgres

import (
	"classconnect-api/pkg/client/postgresql"
	"context"
	"log/slog"
)

const (
	layerScheduleRepository = "repository.schedule."
)

type ScheduleRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewScheduleRepository(client postgresql.Client, logger *slog.Logger) *GroupRepository {
	return &GroupRepository{
		db:     client,
		logger: logger,
	}
}

func (s *ScheduleRepository) CreateSchedule(ctx context.Context) {

}
