package controller

import (
	"fmt"
	"golang_side_project_crud_website/models/users"
	"golang_side_project_crud_website/render_templates"
	"log"
	"net/http"
	"path"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pageContent := PageContent{
			PageTitle: "Sign up",
			PageQuery: nil,
		}
		index := path.Join("templates/users", "create.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
	}else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { log.Fatal(err)}

		uid := users.CreateUser(
			r.Form["userName"][0], r.Form["userEmail"][0], r.Form["userPwd"][0])
		fmt.Println(uid)
		// TODO: get the uid and store in session when user create post add the uid to it
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}