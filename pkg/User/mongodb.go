package users

import (
	"context"
	"fmt"
	"log"

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

func (r *repo) CreateUser(ctx context.Context, user *User) (err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	insertResultID, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResultID)
	return err
}

func (r *repo) GetUser(ctx context.Context, id string) (u *User, err error) {
	panic("not implemented")
}

func (r *repo) ModifyUser(ctx context.Context, id string, user *User) (u *User, err error) {
	panic("not implemented")
}

func (r *repo) DeleteUser(ctx context.Context, id string) (success bool, err error) {
	panic("not implemented")
}
