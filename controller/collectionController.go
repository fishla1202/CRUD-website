package controller

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models"
	"golang_side_project_crud_website/render_templates"
	"net/http"
	"path"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	isUser := config.CheckSessionCookie(w, r)

	if !isUser {
		http.Redirect(w, r, "/user/login/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Form["title"] == nil || r.Form["description"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			collection := models.Collection{
				Title:       r.Form["title"][0],
				Description: r.Form["description"][0],
			}
			err := models.CreateCollection(&collection)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Method == "GET" {

		pageContent := PageContent{
			PageTitle: "Create Collection",
			CsrfTag:   csrf.TemplateField(r),
			IsUser:    isUser,
		}
		index := path.Join("templates/collections", "create.html")
		render_templates.ReturnRenderTemplate(w, index, &pageContent)
		return
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func SearchCollectionPosts(w http.ResponseWriter, r *http.Request) {
	isUser := config.CheckSessionCookie(w, r)
	params := mux.Vars(r)
	collectionTitle := params["collectionTitle"]
	posts, err := models.FindPostByCollectionTitle(collectionTitle)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageContent := newPageContent()
	pageContent.PageTitle = "Discuss"
	pageContent.PageQuery = posts
	pageContent.IsUser = isUser

	index := path.Join("templates", "index.html")
	render_templates.ReturnRenderTemplate(w, index, &pageContent)
	return
}
