package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

type muxRouter struct {
}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

// @Summary Get all the documents inside of deliveries collection.
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}
func (*muxRouter) SERVE(port string) {
	log.Printf("Mux HTTP server running on port: %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}
func (*muxRouter) SWAGGER(uri string) {
	muxDispatcher.PathPrefix(uri).Handler(httpSwagger.WrapHandler)
}
