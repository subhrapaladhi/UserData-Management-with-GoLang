package users

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

type Service interface {
	Register(ctx context.Context, email, name, phone, password string) (interface{}, error)

	Login(ctx context.Context, email, password string) (*User, error)

	GetUserProfile(ctx context.Context, id string) (interface{}, error)

	ModifyUserProfile(ctx context.Context, id, email, name, phone, password string) (interface{}, error)

	DeleteUserProfile(ctx context.Context, id string) (interface{}, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Register(ctx context.Context, email, name, phone, password string) (result interface{}, err error) {
	hasher := md5.New()
	hasher.Write([]byte(password))

	newUser := User{Email: email, Name: name, Phone: phone, Password: hex.EncodeToString(hasher.Sum(nil))}

	return s.repo.CreateUser(ctx, &newUser)
}

func (s *service) Login(ctx context.Context, email, password string) (u *User, err error) {
	panic("not implemented")
}

func (s *service) GetUserProfile(ctx context.Context, id string) (u interface{}, err error) {
	return s.repo.GetUser(ctx, id)
}

func (s *service) ModifyUserProfile(ctx context.Context, id, email, name, phone, password string) (u interface{}, err error) {
	userData := User{}

	if email != "" {
		userData.Email = email
	}
	if name != "" {
		userData.Name = name
	}
	if phone != "" {
		userData.Phone = phone
	}
	if password != "" {
		hasher := md5.New()
		hasher.Write([]byte(password))
		userData.Password = hex.EncodeToString(hasher.Sum(nil))
	}

	return s.repo.ModifyUser(ctx, id, &userData)
}

func (s *service) DeleteUserProfile(ctx context.Context, id string) (u interface{}, err error) {
	return s.repo.DeleteUser(ctx, id)
}
