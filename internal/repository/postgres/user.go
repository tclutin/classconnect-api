package postgres

import (
	"classconnect-api/internal/domain/auth"
	"classconnect-api/pkg/client/postgresql"
	"context"
	"log/slog"
)

const (
	layerUserRepository = "repository.user."
)

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
	sql := `INSERT INTO public.users (username, email, hashed_password, created_at) VALUES ($1, $2, $3, $4)`

	_, err := u.db.Exec(ctx, sql, user.Username, user.Email, user.PasswordHash, user.CreatedAt)

	return err
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (auth.User, error) {
	sql := `SELECT * FROM public.users WHERE username = $1`

	u.logger.Info(layerUserRepository+"GetUserByUsername", slog.String("sql", sql))

	var user auth.User

	row := u.db.QueryRow(ctx, sql, username)

	err := row.Scan(
		&user.ID,
		&user.GroupID,
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

func (u *UserRepository) UpdateUser(ctx context.Context, user auth.User) error {
	sql := `UPDATE public.users SET group_id = $1,
                        username = $2,
                        email = $3,
                        hashed_password = $4,
                        is_banned = $5 WHERE id = $6`

	_, err := u.db.Exec(ctx, sql,
		user.GroupID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.IsBanned,
		user.ID)

	return err
}
