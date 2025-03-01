package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserDB interface {
	Save(context.Context, User) error
	Read(context.Context, string) (User, error)
	Close(context.Context)
}
