package controller

import (
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models"
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)


func IndexHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	allPosts, err := models.FindAllPosts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageContent := newPageContent()
	pageContent.PageTitle = "Discuss"
	pageContent.PageQuery = allPosts
	pageContent.IsUser = isUser

	index := path.Join("templates", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
	return

}
