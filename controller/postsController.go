package controller

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)


func PostDetail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)
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
	return
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError)}

		if r.Form["content"] == nil || r.Form["title"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else {
			session, _ := config.Store.Get(r, "user-info")
			userId := session.Values["userId"]

			post := posts.Post{
				Title:   r.Form["title"][0],
				Content: r.Form["content"][0],
				UserID:  userId.(uint),
			}
			posts.InsertPost(&post)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}else if r.Method == "GET" {

		pageContent := PageContent{
			PageTitle: "Create Post",
			CsrfTag: csrf.TemplateField(r),
			IsUser: isUser,
		}
		index := path.Join("templates/posts", "create.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
		return
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
		return
	}

	session, _ := config.Store.Get(r, "user-info")
	userId := session.Values["userId"]

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError)}
		if r.Form["content"] == nil || r.Form["title"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}else {
			id := r.Form["id"][0]
			postBelongToUserID := posts.FindById(id).Value.(posts.Post).UserID

			if postBelongToUserID != userId {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			title := r.Form["title"][0]
			content := r.Form["content"][0]

			posts.UpdateById(id, title, content)
			http.Redirect(w, r, "/post/edit/" + id, http.StatusSeeOther)
			return
		}
	}else if r.Method == "GET" {
		params := mux.Vars(r)
		id := params["id"]
		post := posts.FindById(id)

		if post.Value.(*posts.Post).UserID != userId {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		pageContent := PageContent{
			PageTitle: "Edit Post",
			PageQuery: post,
			CsrfTag: csrf.TemplateField(r)}

		index := path.Join("templates/posts", "edit.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
		return
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	isUser := config.CheckSessionCookie(w, r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
		return
	}

	session, _ := config.Store.Get(r, "user-info")
	userId := session.Values["userId"]

	params := mux.Vars(r)
	id := params["id"]

	post := posts.FindById(id)

	if post.Value.(*posts.Post).UserID != userId {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	posts.DeleteById(id)
	http.Redirect(w, r, "/user/dashboard/", http.StatusSeeOther)
	return
}