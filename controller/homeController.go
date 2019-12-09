package controller

import (
	"github.com/jinzhu/gorm"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/render_templates"
	//"github.com/gorilla/sessions"
	"log"
	"net/http"
	"path"
)

type PageContent struct {
	PageTitle string
	PageQuery *gorm.DB
}

// TODO: 將login firebase 拿到的uid 存在session裡
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		allPosts := posts.FindAllPosts()

		pageContent := PageContent{
			PageTitle: "FishLa",
			PageQuery: allPosts,
		}

		index := path.Join("templates", "index.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {log.Fatal(err)}
		log.Fatal(r.Form)
	}

}
