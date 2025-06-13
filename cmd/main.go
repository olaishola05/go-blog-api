package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/olaishola05/blog-apis/pkg/config"
	"github.com/olaishola05/blog-apis/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	cfgs, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(map[string]string{"data": "Welcome to Go backend Blogging APIS"})
	}).Methods("GET")

	v1 := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterUserRoutes(v1)
	routes.RegistPostRoutes(v1)

	fmt.Println("Server running on PORT:", cfgs.Port)
	log.Fatal(http.ListenAndServe("localhost:"+cfgs.Port, router))
}
