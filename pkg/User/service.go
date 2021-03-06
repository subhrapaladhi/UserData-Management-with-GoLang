package users

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

type Service interface {
	Register(ctx context.Context, email, name, phone, password string) (string, error)

	Login(ctx context.Context, email, password string) (*User, error)

	GetProfile(ctx context.Context, id string) (*User, error)

	GetAll(ctx context.Context) ([]User, error)

	ModifyProfile(ctx context.Context, id, email, name, phone, password string) (*User, error)

	DeleteProfile(ctx context.Context, id string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Register(ctx context.Context, email, name, phone, password string) (id string, err error) {
	hasher := md5.New()
	hasher.Write([]byte(password))

	newUser := User{Email: email, Name: name, Phone: phone, Password: hex.EncodeToString(hasher.Sum(nil))}
	return s.repo.CreateUser(ctx, &newUser)
}

func (s *service) Login(ctx context.Context, email, password string) (u *User, err error) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return s.repo.GetUserByEmailPassword(ctx, email, hex.EncodeToString(hasher.Sum(nil)))
}

func (s *service) GetAll(ctx context.Context) (u []User, err error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *service) GetProfile(ctx context.Context, id string) (u *User, err error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *service) ModifyProfile(ctx context.Context, id, email, name, phone, password string) (u *User, err error) {
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

func (s *service) DeleteProfile(ctx context.Context, id string) (u *User, err error) {
	return s.repo.DeleteUser(ctx, id)
}
