package users

import (
	"context"
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

func (r *repo) CreateUser(ctx context.Context, user *User) (id string, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	insertResultID, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	id = insertResultID.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *repo) GetUserById(ctx context.Context, id string) (u *User, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	result := User{}
	if err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return &result, err
}

func (r *repo) GetUserByEmailPassword(ctx context.Context, email, password string) (u *User, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	result := User{}
	if err = collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return &result, err
}

func (r *repo) GetAllUsers(ctx context.Context) (interface{}, error) {
	panic("not implemented")
}

func (r *repo) ModifyUser(ctx context.Context, id string, user *User) (u *User, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	// var result interface{}
	result := User{}
	if err = collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, bson.M{"$set": user}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return &result, err
}

func (r *repo) DeleteUser(ctx context.Context, id string) (u *User, err error) {
	collection := r.DB.Database("usermgt").Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	deletedUser := User{}
	if err = collection.FindOneAndDelete(ctx, bson.M{"_id": oid}).Decode(&deletedUser); err != nil {
		log.Fatal(err)
	}
	return &deletedUser, err
}
