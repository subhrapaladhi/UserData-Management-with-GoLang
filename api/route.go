package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/controller"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
	users "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/User"
)

func UserRoutes(mux *http.ServeMux, svc users.Service) {
	mux.HandleFunc("/user/ping", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		if r.Method == http.MethodGet {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(views.ResponseStruct{
				Code: http.StatusOK,
				Data: "user pong",
			})
		}
	})

	mux.Handle("/user/register", controller.RegisterUser(svc))
	mux.Handle("/user/", controller.UserFunctions(svc))
}
