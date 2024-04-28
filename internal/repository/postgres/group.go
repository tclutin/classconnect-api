package postgres

import (
	"classconnect-api/internal/domain/group"
	"classconnect-api/pkg/client/postgresql"
	"context"
	"log/slog"
)

const (
	layerGroupRepository = "repository.user."
)

type GroupRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewGroupRepository(client postgresql.Client, logger *slog.Logger) *GroupRepository {
	return &GroupRepository{
		db:     client,
		logger: logger,
	}
}

func (g *GroupRepository) CreateGroup(ctx context.Context, group group.Group) error {
	sql := `INSERT INTO public.groups (name, code, created_at) VALUES ($1, $2, $3)`

	_, err := g.db.Exec(ctx, sql, group.Name, group.Code, group.CreatedAt)

	return err
}

func (g *GroupRepository) GetGroupByName(ctx context.Context, name string) (group.Group, error) {
	sql := `SELECT * FROM public.groups WHERE name = $1`

	var getGroup group.Group

	row := g.db.QueryRow(ctx, sql, name)

	err := row.Scan(
		&getGroup.ID,
		&getGroup.Name,
		&getGroup.Code,
		&getGroup.MembersCount,
		&getGroup.CreatedAt)

	if err != nil {
		return group.Group{}, err
	}

	return getGroup, nil
}

func (g *GroupRepository) GetAllGroups(ctx context.Context) ([]group.Group, error) {
	sql := `SELECT id, name, members_count, created_at FROM public.groups`

	rows, err := g.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	groups := make([]group.Group, 0)

	for rows.Next() {
		var getGroup group.Group

		err = rows.Scan(
			&getGroup.ID,
			&getGroup.Name,
			&getGroup.MembersCount,
			&getGroup.CreatedAt)

		if err != nil {
			return nil, err
		}

		groups = append(groups, getGroup)
	}

	return groups, nil
}
