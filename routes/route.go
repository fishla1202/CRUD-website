package routes

import (
	"golang_side_project_crud_website/controller"
	"net/http"
)

func Main() {
	http.HandleFunc("/", controller.IndexHandle)

	// post
	http.HandleFunc("/create-post", controller.CreatePost)
	http.HandleFunc("/post-list", controller.PostIndex)
}

