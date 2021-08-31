package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
	users "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/User"
)

func RegisterUser(svc users.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var newUser users.User
			if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
				log.Fatal(err)
			}

			result, err := svc.Register(r.Context(), newUser.Email, newUser.Name, newUser.Phone, newUser.Password)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusCreated)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusCreated,
				Data: result,
			})
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	})
}

func UserFunctions(svc users.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Path[6:]
			user, err := svc.GetUserProfile(context.TODO(), id)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: user,
			})
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	})
}
