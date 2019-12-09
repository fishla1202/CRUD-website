package routes

import (
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/controller"
	"net/http"
	"path"
)

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.IndexHandle).Methods("GET").Name("home")

	// post
	p := r.PathPrefix("/post").Subrouter()
	p.HandleFunc("/create/", controller.CreatePost).Methods("GET", "POST").Name("createPost")
	p.HandleFunc("/{id:[0-9]+}/", controller.PostDetail).Methods("GET").Name("postDetail")
	p.HandleFunc("/delete/{id:[0-9]+}/", controller.DeletePost).Methods("GET").Name("deletePost")
	p.HandleFunc("/edit/{id:[0-9]+}/", controller.EditPost).Methods("GET", "POST").Name("editPost")

	// user
	u := r.PathPrefix("/user").Subrouter()
	u.HandleFunc("/sign-up/", controller.CreateUser).Methods("GET", "POST").Name("createUser")
	u.HandleFunc("/sign-in/", controller.LoginUser).Name("loginUser")

	// load the static file
	r.HandleFunc("/public/firebase_config.js", SendJqueryJs).Methods("GET").Name("firebaseConfig")

	http.Handle("/", r)
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path.Join("static", "firebase_config.js"))
}
