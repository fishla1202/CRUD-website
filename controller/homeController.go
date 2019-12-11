package controller

import (
	"github.com/jinzhu/gorm"
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
}

// TODO: 將login firebase 拿到的uid 存在session裡
func IndexHandle(w http.ResponseWriter, r *http.Request) {

	//cookie, err := r.Cookie("session")
	//if cookie != nil {log.Fatal(123)}
	//if err != nil {log.Fatal(err)}
	allPosts := posts.FindAllPosts()

	pageContent := PageContent{
		PageTitle: "FishLa",
		PageQuery: allPosts,
	}

	index := path.Join("templates", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)


}
