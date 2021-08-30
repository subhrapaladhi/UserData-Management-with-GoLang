package users

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error

	GetUser(ctx context.Context, id string) (*User, error)

	ModifyUser(ctx context.Context, id string, user *User) (*User, error)

	DeleteUser(ctx context.Context, id string) (bool, error)
}
