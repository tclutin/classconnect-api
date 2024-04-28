package postgres

import (
	"classconnect-api/pkg/client/postgresql"
	"context"
	"log/slog"
)

const (
	layerSubscriberRepository = "repository.user."
)

type SubscriberRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewSubscriberRepository(client postgresql.Client, logger *slog.Logger) *SubscriberRepository {
	return &SubscriberRepository{
		db:     client,
		logger: logger,
	}
}

func (s *SubscriberRepository) CreateTelegramSubscriber(ctx context.Context) {
	sql := `INSERT INTO public.su`
}

func (s *SubscriberRepository) CreateMobileSubscriber(ctx context.Context) {

}
