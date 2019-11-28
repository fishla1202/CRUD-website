package controller

import (
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/render_templates"
	"log"
	"net/http"
	"path"
)

// TODO: post list page
func PostIndex(w http.ResponseWriter, r *http.Request) {
	allPosts := posts.FindAllPosts()

	pageContent := PageContent{
		PageTitle: "Post List",
		PageQuery: allPosts,
	}

	index := path.Join("templates/posts", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { log.Fatal(err)}
		if r.Form["content"] == nil || r.Form["title"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else {
			title := r.Form["title"][0]
			content := r.Form["content"][0]
			posts.InsertPost(title, content)
			http.Redirect(w, r, "/post-list", http.StatusSeeOther)
		}
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
