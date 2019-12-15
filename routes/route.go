package routes

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/controller"
	"log"
	"net/http"
	"os"
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
	u.HandleFunc("/login/", controller.LoginUser).Methods("GET").Name("loginUser")
	u.HandleFunc("/login/", config.SetLoginSession).Methods("POST").Name("setLoginSession")

	// load the static file
	r.HandleFunc("/public/firebase_config.js", SendJqueryJs).Methods("GET").Name("firebaseConfig")

	csrfKey := make([]byte, 32)
	env := os.Getenv("APP")

	var csrfMiddleware func(http.Handler) http.Handler

	if env == "dev" {
		csrfMiddleware = csrf.Protect(csrfKey, csrf.Secure(false))
	}else if  env == "production"{
		csrfMiddleware = csrf.Protect(csrfKey)
	}else {
		log.Fatal("env app setup error please select type production or dev")
	}
	// gorilla csrf token 預設只對同一個route or 子 route有效 https://github.com/gorilla/csrf/issues/32 研究好久才發現
	http.Handle("/", csrfMiddleware(r))
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path.Join("static", "firebase_config.js"))
}
