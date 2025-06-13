package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/olaishola05/blog-apis/pkg/models"
	"github.com/olaishola05/blog-apis/pkg/utils"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	posts, err := models.GetPosts(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if posts == nil {
		posts = []models.Post{}
	}

	response := utils.UserResponse{
		Status:  http.StatusOK,
		Message: "Posts retrieved successfully",
		Data:    posts,
	}
	res, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {

}

func GetPost(w http.ResponseWriter, r *http.Request) {

}

func UpdatePostById(w http.ResponseWriter, r *http.Request) {

}

func DeletePostById(w http.ResponseWriter, r *http.Request) {

}
