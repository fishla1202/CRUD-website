package routes

import (
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/controller"
	"net/http"
)

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.IndexHandle).Methods("GET")

	// post
	r.HandleFunc("/create-post", controller.CreatePost).Methods("POST")
	r.HandleFunc("/post-list", controller.PostIndex).Methods("GET")
	r.HandleFunc("/post/{id}", controller.PostDetail).Methods("GET")

	http.Handle("/", r)
}

