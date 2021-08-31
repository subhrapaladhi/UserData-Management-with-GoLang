package users

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *repo) CreateUser(ctx context.Context, user *User) (result interface{}, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	insertResultID, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(insertResultID)
	return insertResultID, nil
}

func (r *repo) GetUser(ctx context.Context, id string) (u interface{}, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	result := User{}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	if err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	// res := collection.FindOne(ctx, bson.M{"_id": oid})
	// res.Decode(&result)
	// fmt.Println(result)
	return result, err
}

func (r *repo) ModifyUser(ctx context.Context, id string, user *User) (u *User, err error) {
	panic("not implemented")
}

func (r *repo) DeleteUser(ctx context.Context, id string) (success bool, err error) {
	panic("not implemented")
}
