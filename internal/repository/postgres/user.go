package postgres

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/pkg/client/postgresql"
	"context"
	"log/slog"
)

const layer_user_repository = "repository.user."

type UserRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewUserRepository(client postgresql.Client, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		db:     client,
		logger: logger,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, user auth.User) error {
	sql := "INSERT INTO public.user (username, email, hashed_password, created_at) VALUES ($1, $2, $3, $4)"

	u.logger.Info(layer_user_repository+"CreateUser", slog.String("sql", sql))

	_, err := u.db.Exec(ctx, sql, user.Username, user.Email, user.PasswordHash, user.CreatedAt)

	return err
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (auth.User, error) {
	sql := "SELECT * FROM public.user WHERE username = $1"

	u.logger.Info(layer_user_repository+"GetUserByUsername", slog.String("sql", sql))

	var user auth.User

	row := u.db.QueryRow(ctx, sql, username)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.IsBanned,
		&user.CreatedAt)

	if err != nil {
		return auth.User{}, err
	}

	return user, nil
}
