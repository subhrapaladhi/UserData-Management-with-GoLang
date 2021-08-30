package controller

import (
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

			err := svc.Register(r.Context(), newUser.Email, newUser.Name, newUser.Phone, newUser.Password)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusCreated)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusCreated,
				Data: newUser,
			})
		}
	})
}
