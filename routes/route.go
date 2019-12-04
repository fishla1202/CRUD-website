package routes

import (
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/controller"
	"net/http"
	"path"
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
	r.HandleFunc("/user/sign-up/", controller.CreateUser)

	// load the static file
	r.HandleFunc("/public/firebase_config.js", SendJqueryJs)

	http.Handle("/", r)
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	/*TODO: add static dir and add firebase_config.js file*/
	http.ServeFile(w, r, path.Join("static", "firebase_config.js"))
}
