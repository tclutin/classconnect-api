package group

import (
	"errors"
	"time"
)

var (
	ErrNotFound      = errors.New("group not found")
	ErrWrongCode     = errors.New("wrong group code")
	ErrAlreadyExists = errors.New("group already exists with this name")
)

type Group struct {
	ID           uint64
	Name         string
	Code         string
	MembersCount uint
	CreatedAt    time.Time
}
