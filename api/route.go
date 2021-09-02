package api

import (
	"encoding/json"
	"net/http"

	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/controller"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/middleware"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
	users "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/User"
)

func UserRoutes(mux *http.ServeMux, svc users.Service) {
	mux.HandleFunc("/user/ping", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: "user pong",
			})
		}
	})

	mux.Handle("/user/register", controller.RegisterUser(svc))
	mux.Handle("/user/login", controller.LoginUser(svc))
	mux.Handle("/user/", middleware.Validate(controller.UserFunctions(svc), []string{"user", "admin"}))
}
