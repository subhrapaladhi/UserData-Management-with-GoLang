package admins

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

type Service interface {
	Register(ctx context.Context, email, name, phone, password string) (string, error)

	Login(ctx context.Context, email, password string) (*Admin, error)

	GetProfile(ctx context.Context, id string) (*Admin, error)

	GetAll(ctx context.Context) ([]Admin, error)

	ModifyProfile(ctx context.Context, id, email, name, phone, password string) (*Admin, error)

	DeleteProfile(ctx context.Context, id string) (*Admin, error)
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

	newAdmin := Admin{Email: email, Name: name, Phone: phone, Password: hex.EncodeToString(hasher.Sum(nil))}
	return s.repo.CreateAdmin(ctx, &newAdmin)
}

func (s *service) Login(ctx context.Context, email, password string) (admin *Admin, err error) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return s.repo.GetAdminByEmailPassword(ctx, email, hex.EncodeToString(hasher.Sum(nil)))
}

func (s *service) GetProfile(ctx context.Context, id string) (admin *Admin, err error) {
	return s.repo.GetAdminById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) (adminList []Admin, err error) {
	return s.repo.GetAllAdmins(ctx)
}

func (s *service) ModifyProfile(ctx context.Context, id, email, name, phone, password string) (admin *Admin, err error) {
	adminData := Admin{}

	if email != "" {
		adminData.Email = email
	}
	if name != "" {
		adminData.Name = name
	}
	if phone != "" {
		adminData.Phone = phone
	}
	if password != "" {
		hasher := md5.New()
		hasher.Write([]byte(password))
		adminData.Password = hex.EncodeToString(hasher.Sum(nil))
	}

	return s.repo.ModifyAdmin(ctx, id, &adminData)
}

func (s *service) DeleteProfile(ctx context.Context, id string) (admin *Admin, err error) {
	return s.repo.DeleteAdmin(ctx, id)
}
