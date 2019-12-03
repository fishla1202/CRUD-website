package controller

import (
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	pageContent := PageContent{
		PageTitle: "Sing up",
		PageQuery: nil,
	}
	index := path.Join("templates/users", "create.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
}