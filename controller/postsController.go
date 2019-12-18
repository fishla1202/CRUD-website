package controller

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/models/users"
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)


func PostDetail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(r)
	params := mux.Vars(r)
	id := params["id"]
	post := posts.FindById(id)

	// 當回傳是interface時 需要定義回傳是什麼值才能提取裡面的屬性
	pageContent := PageContent{
		PageTitle: post.Value.(*posts.Post).Title,
		PageQuery: post,
		IsUser: isUser,
	}

	index := path.Join("templates/posts", "detail.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError)}
		if r.Form["content"] == nil || r.Form["title"] == nil || r.Form["uid"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else {
			uid, err := r.Cookie("userInfo")
			if err != nil {http.Redirect(w, r, "/uesr/login/", http.StatusSeeOther)}
			userId := users.FindUserByUID(uid.Value)

			post := posts.Post{
				Title:   r.Form["title"][0],
				Content: r.Form["content"][0],
				UserID:  userId,
			}
			posts.InsertPost(&post)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}else if r.Method == "GET" {

		pageContent := PageContent{
			PageTitle: "Create Post",
			CsrfTag: csrf.TemplateField(r),
			IsUser: isUser,
		}
		index := path.Join("templates/posts", "create.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError)}
		if r.Form["content"] == nil || r.Form["title"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else {
			uid, err := r.Cookie("userInfo")
			if err != nil {http.Redirect(w, r, "/uesr/login/", http.StatusSeeOther)}
			userId := users.FindUserByUID(uid.Value)
			title := r.Form["title"][0]
			content := r.Form["content"][0]
			id := r.Form["id"][0]
			postBelongToUserID := posts.FindById(id).Value.(posts.Post).UserID

			if postBelongToUserID != userId {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
			
			posts.UpdateById(id, title, content)
			http.Redirect(w, r, "/post/edit/" + id, http.StatusSeeOther)
		}
	}else if r.Method == "GET" {
		params := mux.Vars(r)
		id := params["id"]
		post := posts.FindById(id)

		pageContent := PageContent{
			PageTitle: "Edit Post",
			PageQuery: post,
			CsrfTag: csrf.TemplateField(r)}

		index := path.Join("templates/posts", "edit.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	id := params["id"]
	//fmt.Println(id)
	posts.DeleteById(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}