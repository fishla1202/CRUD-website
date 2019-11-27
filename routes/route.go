package routes

import (
	"golang_side_project_crud_website/controller"
	"net/http"
)

func Main() {
	http.HandleFunc("/", controller.IndexHandle)
	http.HandleFunc("/create-post", controller.CreatePost)
}

