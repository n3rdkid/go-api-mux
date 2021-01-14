package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

// NewMuxRouter -> Return a muxRouter
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")

}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP router listening on port : %s", port)
	http.ListenAndServe(port, muxDispatcher)
}
