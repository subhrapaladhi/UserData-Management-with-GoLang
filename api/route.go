package api

import (
	"encoding/json"
	"net/http"

	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api/views"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.ResponseStruct{
				Code: http.StatusOK,
				Body: "pong",
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(data)
		}
	})

	return mux
}
