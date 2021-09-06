package api

import (
	"net/http"

	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/controller"
	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/middleware"
	admins "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/Admin"
	users "github.com/subhrapaladhi/User-Data-Management-with-GoLang/pkg/User"
)

func UserRoutes(mux *http.ServeMux, svc users.Service) {
	mux.Handle("/user/register", controller.RegisterUser(svc))
	mux.Handle("/user/login", controller.LoginUser(svc))
	mux.Handle("/user/", middleware.Validate(controller.UserFunctions(svc), []string{"user", "admin"}))
}

func AdminRoutes(mux *http.ServeMux, adminSvc admins.Service, userSvc users.Service) {
	mux.Handle("/admin/register", controller.RegisterAdmin(adminSvc))
	mux.Handle("/admin/login", controller.LoginAdmin(adminSvc))
	mux.Handle("/admin/", middleware.Validate(controller.AdminFunctions(adminSvc), []string{"admin"}))
	mux.Handle("/usermgt/", middleware.Validate(controller.UserMgtFunc(userSvc), []string{"admin"}))
}
