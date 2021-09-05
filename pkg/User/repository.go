package users

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) (string, error)

	GetUserById(ctx context.Context, id string) (*User, error)

	GetUserByEmailPassword(ctx context.Context, email, password string) (*User, error)

	GetAllUsers(ctx context.Context) ([]User, error)

	ModifyUser(ctx context.Context, id string, user *User) (*User, error)

	DeleteUser(ctx context.Context, id string) (*User, error)
}
