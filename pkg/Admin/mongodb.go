package admins

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	DB *mongo.Client
}

func NewMongodbRepo(db *mongo.Client) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateAdmin(ctx context.Context, admin *Admin) (id string, err error) {
	panic("not implemented")
}

func (r *repo) GetAdminById(ctx context.Context, id string) (admin *Admin, err error) {
	panic("not implemented")
}

func (r *repo) GetAdminByEmailPassword(ctx context.Context, email, password string) (admin *Admin, err error) {
	panic("not implemented")
}

func (r *repo) ModifyAdmin(ctx context.Context, id string, admin *Admin) (a *Admin, err error) {
	panic("not implemented")
}

func (r *repo) DeleteAdmin(ctx context.Context, id string) (admin *Admin, err error) {
	panic("not implemented")
}
