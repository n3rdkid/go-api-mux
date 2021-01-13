package main

import (
	"encoding/json"
	"go-api-mux/entity"
	"go-api-mux/repository"
	"math/rand"
	"net/http"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Failed to load posts!}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}
func addPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Oops something went wrong!}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
