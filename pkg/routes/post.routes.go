package routes

import (
	"github.com/gorilla/mux"
	"github.com/olaishola05/blog-apis/pkg/controllers"
)

func RegistPostRoutes(router *mux.Router) {
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", controllers.CreateNewPost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.UpdatePostById).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePostById).Methods("DELETE")
}
