package controller

import (
	"fmt"
	"github.com/gorilla/csrf"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/models/users"
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)


func UserIndex(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	isUser := config.CheckSessionCookie(w, r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
	}else {
		session, _ := config.Store.Get(r, "user-info")
		userID := session.Values["userId"]
		qs := posts.FindByUserId(userID.(uint))

		pageContent := PageContent{
			PageTitle: "Dashboard",
			PageQuery: qs,
			//CsrfTag: csrf.TemplateField(r),
			IsUser: isUser,
		}

		index := path.Join("templates/users", "index.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}

}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	isUser := config.CheckSessionCookie(w, r)

	if isUser {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == "GET" {
		pageContent := PageContent{
			PageTitle: "Sign up",
			PageQuery: nil,
			CsrfTag: csrf.TemplateField(r),
			IsUser: isUser,
		}

		index := path.Join("templates/users", "create.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) }

	// TODO: 判斷使用者重複註冊問題
		uid := users.CreateUser(
			r.Form["userName"][0], r.Form["userEmail"][0], r.Form["userPwd"][0])
		fmt.Println(uid)
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	if isUser {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	pageContent := PageContent{
		PageTitle: "User login",
		PageQuery: nil,
		CsrfTag: csrf.TemplateField(r),
		IsUser: isUser,
	}

	index := path.Join("templates/users", "login.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)

}