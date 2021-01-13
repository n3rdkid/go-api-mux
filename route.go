package main

import (
	"encoding/json"
	"net/http"
)

// Post -> Post Model
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		Post{ID: 1, Title: "Title 1", Text: "Text 1"},
		Post{ID: 2, Title: "Title 2", Text: "Text 2"},
		Post{ID: 3, Title: "Title 3", Text: "Text 3"},
	}
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Oops something went wrong!}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

}
func addPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Oops something went wrong!}`))
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Oops something went wrong!}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
