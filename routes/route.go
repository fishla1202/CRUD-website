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
	r.HandleFunc("/create-post/", controller.CreatePost)
	r.HandleFunc("/post/{id:[0-9]+}/", controller.PostDetail).Methods("GET")
	r.HandleFunc("/post/delete/{id:[0-9]+}/", controller.DeletePost).Methods("GET")
	r.HandleFunc("/post/edit/{id:[0-9]+}/", controller.EditPost)

	// user
	r.HandleFunc("/user/sing-up/", controller.CreateUser)

	http.Handle("/", r)
}

