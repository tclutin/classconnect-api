package auth

import (
	"errors"
	"time"
)

var (
	ErrAlreadyExist = errors.New("user already exist")
	ErrNotFound     = errors.New("user not found")
)

type User struct {
	ID           uint64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
