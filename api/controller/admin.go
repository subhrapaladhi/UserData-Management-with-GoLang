package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
	admins "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/Admin"
)

func RegisterAdmin(svc admins.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		var newAdmin admins.Admin
		if err := json.NewDecoder(r.Body).Decode(&newAdmin); err != nil {
			log.Fatal(err)
		}

		id, err := svc.Register(r.Context(), newAdmin.Email, newAdmin.Name, newAdmin.Phone, newAdmin.Password)
		if err != nil {
			log.Fatal(err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   id,
			"role": "admin",
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

func LoginAdmin(svc admins.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		var adminData admins.Admin
		if err := json.NewDecoder(r.Body).Decode(&adminData); err != nil {
			log.Fatal(err)
		}

		admin, err := svc.Login(context.TODO(), adminData.Email, adminData.Password)
		if err != nil {
			log.Fatal(err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   admin.Id,
			"role": "admin",
		})
		tokenString, err := token.SignedString([]byte(viper.GetString("JWTSECRET")))
		if err != nil {
			log.Fatal(err)
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"data":  admin,
			"token": tokenString,
		})
	})
}

func AdminFunctions(svc admins.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// uid := r.Context().Value("jwtData").(jwt.MapClaims)["id"].(string) //getting id from the jwt data set in the context by the auth middleware
		if r.Method == http.MethodGet { // Get using id
			id := r.URL.Path[7:]
			fmt.Println(id)
			admin, err := svc.GetProfile(context.TODO(), id)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: admin,
			})
		} else if r.Method == http.MethodPatch { // Edit user data
			id := r.URL.Path[7:]
			var adminData admins.Admin
			if err := json.NewDecoder(r.Body).Decode(&adminData); err != nil {
				log.Fatal(err)
			}
			result, err := svc.ModifyProfile(context.TODO(), id, adminData.Email, adminData.Name, adminData.Phone, adminData.Password)
			if err != nil {
				log.Fatal(err)
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: result,
			})
		} else if r.Method == http.MethodDelete { // Delete user
			id := r.URL.Path[7:]
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
