package controller

import (
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models"
	"golang_side_project_crud_website/render_templates"
	"html/template"
	"net/http"
	"path"
)

type PageContent struct {
	PageTitle string
	PageQuery interface{}
	CsrfTag template.HTML
	IsUser bool
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	allPosts, err := models.FindAllPosts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageContent := PageContent{
		PageTitle: "Discuss",
		PageQuery: allPosts,
		IsUser: isUser,
	}

	index := path.Join("templates", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
	return

}
