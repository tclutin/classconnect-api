package postgres

import (
	"classconnect-api/internal/domain/schedule"
	"classconnect-api/pkg/client/postgresql"
	"context"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

const (
	layerScheduleRepository = "repository.schedule."
)

type ScheduleRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewScheduleRepository(client postgresql.Client, logger *slog.Logger) *ScheduleRepository {
	return &ScheduleRepository{
		db:     client,
		logger: logger,
	}
}

func (s *ScheduleRepository) CreateSchedule(ctx context.Context, schedule schedule.UploadScheduleDTO, groupID *uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, week := range schedule.Weeks {
		weekID, err := s.CreateWeek(ctx, tx, week, groupID)
		if err != nil {
			return err
		}

		for _, day := range week.Days {
			dayID, err := s.CreateDay(ctx, tx, weekID, day)
			if err != nil {
				return err
			}

			for _, subject := range day.Subjects {
				err := s.CreateSubject(ctx, tx, dayID, subject)
				if err != nil {
					return err
				}
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) CreateWeek(ctx context.Context, tx pgx.Tx, week schedule.WeekDTO, groupID *uint64) (uint64, error) {
	sql := `INSERT INTO public.weeks (group_id, is_even) VALUES ($1, $2) RETURNING id`

	var weekID uint64

	err := tx.QueryRow(ctx, sql, groupID, week.IsEven).Scan(&weekID)
	if err != nil {
		return 0, err
	}

	return weekID, nil
}

func (s *ScheduleRepository) CreateDay(ctx context.Context, tx pgx.Tx, weekID uint64, day schedule.DayDTO) (uint64, error) {
	sql := `INSERT INTO public.days (week_id, day_of_week) VALUES ($1, $2) RETURNING id`

	var dayID uint64

	err := tx.QueryRow(ctx, sql, weekID, day.DayNumber).Scan(&dayID)
	if err != nil {
		return 0, err
	}

	return dayID, nil
}

func (s *ScheduleRepository) CreateSubject(ctx context.Context, tx pgx.Tx, dayID uint64, day schedule.SubjectDTO) error {
	sql := `INSERT INTO public.subjects (day_id,
                             teacher,
                             name,
                             cabinet,
                             description,
                             time_start,
                             time_end) VALUES ($1,$2, $3, $4, $5, $6, $7)`

	_, err := tx.Exec(ctx, sql,
		dayID,
		day.Teacher,
		day.Name,
		day.Cabinet,
		day.Description,
		day.StartTime,
		day.EndTime)

	if err != nil {
		return err
	}

	return nil
}
