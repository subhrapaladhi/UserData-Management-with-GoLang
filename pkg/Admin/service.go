package admins

import "context"

type Service interface {
	// Services to manage the admin user
	Register(ctx context.Context, email, name, phone, password string) (string, error)

	Login(ctx context.Context, email, password string) (*Admin, error)

	GetProfile(ctx context.Context, id string) (*Admin, error)

	ModifyProfile(ctx context.Context, id, email, name, phone, password string) (*Admin, error)

	DeleteProfile(ctx context.Context, id string) (*Admin, error)

	// Services to manager other users
	CreateUser(ctx context.Context, email, name, phone, password string) (string, error)

	GetUserProfile(ctx context.Context, id string) (interface{}, error)

	GetAllUserProfiles(ctx context.Context) (interface{}, error)

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

func (s *service) Register(ctx context.Context, email, name, phone, password string) (id string, err error) {
	panic("not implemented")
}

func (s *service) Login(ctx context.Context, email, password string) (admin *Admin, err error) {
	panic("not implemented")
}

func (s *service) GetProfile(ctx context.Context, id string) (admin *Admin, err error) {
	panic("not implemented")
}

func (s *service) ModifyProfile(ctx context.Context, id, email, name, phone, password string) (admin *Admin, err error) {
	panic("not implemented")
}

func (s *service) DeleteProfile(ctx context.Context, id string) (admin *Admin, err error) {
	panic("not implemented")
}

func (s *service) CreateUser(ctx context.Context, email, name, phone, password string) (id string, err error) {
	panic("not implemented")
}

func (s *service) GetUserProfile(ctx context.Context, id string) (user interface{}, err error) {
	panic("not implemented")
}

func (s *service) GetAllUserProfiles(ctx context.Context) (user interface{}, err error) {
	panic("not implemented")
}

func (s *service) ModifyUserProfile(ctx context.Context, id, email, name, phone, password string) (user interface{}, err error) {
	panic("not implemented")
}

func (s *service) DeleteUserProfile(ctx context.Context, id string) (user interface{}, err error) {
	panic("not implemented")
}
