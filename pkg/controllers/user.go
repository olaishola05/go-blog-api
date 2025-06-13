package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/olaishola05/blog-apis/pkg/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	users, err := models.GetAllUsers(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if users == nil {
		users = []models.User{}
	}

	res, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

}
