package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
	admins "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/Admin"
	users "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/User"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() (*mongo.Client, error) {
	// set client options
	DBUSER := viper.GetString("DBUSER")
	DBUSERPASS := viper.GetString("DBUSERPASS")
	DBURL := fmt.Sprintf("mongodb://%s:%s@localhost:27017/?authSource=admin", DBUSER, DBUSERPASS)
	clientOptions := options.Client().ApplyURI(DBURL)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func serverInit(db *mongo.Client) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := users.NewMongodbRepo(db)
	userSvc := users.NewService(userRepo)

	adminRepo := admins.NewMongodbRepo(db)
	adminSvc := admins.NewService(adminRepo)

	mux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: "pong",
			})
		}
	})

	api.UserRoutes(mux, userSvc)
	api.AdminRoutes(mux, adminSvc, userSvc)

	return mux
}

func main() {
	// LOADING ENV VARIABLES
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// CONNECTING MONGODB
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

	// CREATING THE SERVER MUX
	mux := serverInit(client)

	fmt.Println("Server started!")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
