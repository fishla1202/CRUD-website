package controller

import (
	"github.com/jinzhu/gorm"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/render_templates"
	"html/template"
	"net/http"
	"path"
)

type PageContent struct {
	PageTitle string
	PageQuery *gorm.DB
	CsrfTag template.HTML
	IsUser bool
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	allPosts := posts.FindAllPosts()

	pageContent := PageContent{
		PageTitle: "Discuss",
		PageQuery: allPosts,
		IsUser: isUser,
	}

	index := path.Join("templates", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
	return

}
