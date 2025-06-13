package routes

import (
	"github.com/gorilla/mux"
	"github.com/olaishola05/blog-apis/pkg/controllers"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateNewUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserById).Methods("DELETE")
}
