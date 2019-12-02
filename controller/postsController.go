package controller

import (
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/render_templates"
	"log"
	"net/http"
	"path"
)


func PostDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	post := posts.FindById(id)

	// 當回傳是interface時 需要定義回傳是什麼值才能提取裡面的屬性
	pageContent := PageContent{
		PageTitle: post.Value.(*posts.Post).Title,
		PageQuery: post,
	}

	index := path.Join("templates/posts", "detail.html")
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
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}else if r.Method == "GET" {
		pageContent := PageContent{PageTitle: "Create Post"}
		index := path.Join("templates/posts", "create.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditPost(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { log.Fatal(err)}
		if r.Form["content"] == nil || r.Form["title"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else {
			title := r.Form["title"][0]
			content := r.Form["content"][0]
			id := r.Form["id"][0]
			posts.UpdateById(id, title, content)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}else if r.Method == "GET" {
		params := mux.Vars(r)
		id := params["id"]
		post := posts.FindById(id)
		pageContent := PageContent{
			PageTitle: "Edit Post",
			PageQuery: post}
		index := path.Join("templates/posts", "edit.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	//fmt.Println(id)
	posts.DeleteById(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}