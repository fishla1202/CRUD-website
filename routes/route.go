package routes

import (
	"CRUD/controller"
	"net/http"
)

func Main() {
	http.HandleFunc("/", controller.IndexHandle)
	http.HandleFunc("/create-post", controller.CreatePost)
}

