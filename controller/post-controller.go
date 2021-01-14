package controller

import (
	"encoding/json"
	"go-api-mux/entity"
	"go-api-mux/errors"
	"go-api-mux/service"
	"net/http"
)

var (
	postService service.PostService
)

// PostController ->
type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPosts(w http.ResponseWriter, r *http.Request)
}
type controller struct{}

// NewPostController ->
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

// GetPosts -> Return all posts
func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Failed to load posts"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

// AddPosts -> Add new post
func (*controller) AddPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Failed to unmarshal data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Failed to store post "})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
