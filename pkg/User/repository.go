package users

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) (interface{}, error)

	GetUser(ctx context.Context, id string) (interface{}, error)

	ModifyUser(ctx context.Context, id string, user *User) (interface{}, error)

	DeleteUser(ctx context.Context, id string) (bool, error)
}
