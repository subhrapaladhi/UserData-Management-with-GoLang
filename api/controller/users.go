package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
	users "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/User"
)

func RegisterUser(svc users.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		var newUser users.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			log.Fatal(err)
		}

		id, err := svc.Register(r.Context(), newUser.Email, newUser.Name, newUser.Phone, newUser.Password)
		if err != nil {
			log.Fatal(err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   id,
			"role": "user",
		})

		tokenString, err := token.SignedString([]byte(viper.GetString("JWTSECRET")))
		if err != nil {
			log.Fatal(err)
		}
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"data":  id,
			"token": tokenString,
		})
	})
}

func LoginUser(svc users.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		var userData users.User
		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			log.Fatal(err)
		}

		user, err := svc.Login(context.TODO(), userData.Email, userData.Password)
		if err != nil {
			log.Fatal(err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   user.Id,
			"role": "user",
		})
		tokenString, err := token.SignedString([]byte(viper.GetString("JWTSECRET")))
		if err != nil {
			log.Fatal(err)
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"data":  user,
			"token": tokenString,
		})
	})
}

func UserFunctions(svc users.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value("jwtData").(jwt.MapClaims)["id"].(string) //getting id from the jwt data set in the context by the auth middleware
		if r.Method == http.MethodGet {                                    // Get using id
			id := r.URL.Path[6:]
			if uid != id {
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte("Unauthorized: JWT id doesn't match query id"))
				return
			}
			user, err := svc.GetProfile(context.TODO(), id)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: user,
			})
		} else if r.Method == http.MethodPatch { // Edit user data
			id := r.URL.Path[6:]
			if uid != id {
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte("Unauthorized: JWT id doesn't match query id"))
				return
			}
			var userData users.User
			if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
				log.Fatal(err)
			}
			result, err := svc.ModifyProfile(context.TODO(), id, userData.Email, userData.Name, userData.Phone, userData.Password)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: result,
			})
		} else if r.Method == http.MethodDelete { // Delete user
			id := r.URL.Path[6:]
			if uid != id {
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte("Unauthorized: JWT id doesn't match query id"))
				return
			}
			deletedUser, err := svc.DeleteProfile(context.TODO(), id)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: deletedUser,
			})
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	})
}
