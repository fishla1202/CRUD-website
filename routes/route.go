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

	// set user login session
	r.HandleFunc("/sessionLogin/", config.SetLoginSession).Methods("POST").Name("setLoginSession")

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

	http.Handle("/", csrfMiddleware(r))
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path.Join("static", "firebase_config.js"))
}
