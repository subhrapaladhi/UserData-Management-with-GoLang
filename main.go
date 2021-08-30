package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Invalid type assertion")
	}

	return value
}

func connectDB() (*mongo.Client, error) {
	// set client options
	DBUSER := getEnv("DBUSER")
	DBUSERPASS := getEnv("DBUSERPASS")
	DBURL := fmt.Sprintf("mongodb://%s:%s@localhost:27017/?authSource=admin", DBUSER, DBUSERPASS)
	clientOptions := options.Client().ApplyURI(DBURL)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	return client, err
}

func main() {
	client, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	// user1 := users.User{
	// 	Email:    "subhra@gmail.com",
	// 	Name:     "subhra",
	// 	Phone:    "9003814273",
	// 	Password: "secret",
	// }
	// collection := client.Database("usermgt").Collection("users")
	// insertResult, err := collection.InsertOne(context.TODO(), user1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(insertResult)

	mux := api.Register()

	fmt.Println("Server started!")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
