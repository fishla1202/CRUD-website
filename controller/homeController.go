package controller

import (
	"github.com/jinzhu/gorm"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)

type PageContent struct {
	PageTitle string
	PageQuery *gorm.DB
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {

	allPosts := posts.FindAllPosts()

	pageContent := PageContent{
		PageTitle: "FishLa",
		PageQuery: allPosts,
	}

	index := path.Join("templates", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)

}
