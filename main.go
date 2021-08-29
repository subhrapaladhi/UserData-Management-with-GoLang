package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subhrapaladhi/User-Data-Management-with-GoLang/api"
)

func main() {
	mux := api.Register()

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
