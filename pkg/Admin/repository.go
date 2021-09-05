package admins

import "context"

type Repository interface {
	CreateAdmin(ctx context.Context, admin *Admin) (string, error)

	GetAdminById(ctx context.Context, id string) (*Admin, error)

	GetAdminByEmailPassword(ctx context.Context, email, password string) (*Admin, error)

	GetAllAdmins(ctx context.Context) ([]Admin, error)

	ModifyAdmin(ctx context.Context, id string, admin *Admin) (*Admin, error)

	DeleteAdmin(ctx context.Context, id string) (*Admin, error)
}
