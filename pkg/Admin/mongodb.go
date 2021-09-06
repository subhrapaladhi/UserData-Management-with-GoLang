package admins

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

func (r *repo) CreateAdmin(ctx context.Context, admin *Admin) (id string, err error) {
	collection := r.DB.Database("usermgt").Collection("admins")
	insertResultID, err := collection.InsertOne(ctx, admin)
	if err != nil {
		log.Fatal(err)
	}
	id = insertResultID.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *repo) GetAdminById(ctx context.Context, id string) (admin *Admin, err error) {
	collection := r.DB.Database("usermgt").Collection("admins")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	result := Admin{}
	if err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return &result, err
}

func (r *repo) GetAdminByEmailPassword(ctx context.Context, email, password string) (admin *Admin, err error) {
	collection := r.DB.Database("usermgt").Collection("admins")
	result := Admin{}
	if err = collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return &result, err
}

func (r *repo) GetAllAdmins(ctx context.Context) (adminList []Admin, err error) {
	collection := r.DB.Database("usermgt").Collection("admins")
	results, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var temp Admin
	for results.Next(ctx) {
		err = results.Decode(&temp)
		if err != nil {
			log.Fatal(err)
		}
		adminList = append(adminList, temp)
	}

	return adminList, err
}

func (r *repo) ModifyAdmin(ctx context.Context, id string, admin *Admin) (a *Admin, err error) {
	collection := r.DB.Database("usermgt").Collection("admins")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	// var result interface{}
	result := Admin{}
	if err = collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, bson.M{"$set": admin}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return &result, err
}

func (r *repo) DeleteAdmin(ctx context.Context, id string) (admin *Admin, err error) {
	collection := r.DB.Database("usermgt").Collection("admins")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	deletedAdmin := Admin{}
	if err = collection.FindOneAndDelete(ctx, bson.M{"_id": oid}).Decode(&deletedAdmin); err != nil {
		log.Fatal(err)
	}
	return &deletedAdmin, err
}
