package users

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	result := User{}
	if err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result, err
}

func (r *repo) ModifyUser(ctx context.Context, id string, user *User) (u interface{}, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	var result interface{}
	if err = collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, bson.M{"$set": user}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result, err
}

func (r *repo) DeleteUser(ctx context.Context, id string) (success bool, err error) {
	panic("not implemented")
}
