package main

import (
	"go-api-mux/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	utils.LoadEnv()

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPosts).Methods("POST")
	log.Println("Server listening on port :", os.Getenv("SERVER_PORT"))
	log.Fatalln(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))
}
