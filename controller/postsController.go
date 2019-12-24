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


func PostDetail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)
	params := mux.Vars(r)
	id := params["id"]
	post, err := models.FindPostById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 當回傳是interface時 需要定義回傳是什麼值才能提取裡面的屬性
	pageContent := PageContent{
		PageTitle: post.Title,
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
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Form["content"] == nil ||
			r.Form["title"] == nil ||
			r.Form["collectionID"] == nil{

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		}else {
			session, _ := config.Store.Get(r, "user-info")
			userId := session.Values["userId"]

			collection, err := models.FindCollectionByID(r.Form["collectionID"][0])

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				Title:   r.Form["title"][0],
				Content: r.Form["content"][0],
				UserID: userId.(uint),
				CollectionID: collection.ID,
				Collection: collection,
			}

			err = models.InsertPost(&post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}else if r.Method == "GET" {

		collections, err := models.FindAllCollections()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pageContent := PageContent{
			PageTitle: "Create Post",
			PageQuery: collections,
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
	params := mux.Vars(r)
	id := params["id"]
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError)}

		if r.Form["content"] == nil ||
			r.Form["title"] == nil ||
			r.Form["collectionID"] == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}else {
			post, err := models.FindPostById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if post.UserID != userId {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			title := r.Form["title"][0]
			content := r.Form["content"][0]

			// TODO fix error
			collection, err := models.FindCollectionByID(r.Form["collectionID"][0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var updatePost = map[string]interface{}{}
			updatePost["Collection"] = collection
			updatePost["CollectionID"] = collection.ID
			updatePost["Title"] = title
			updatePost["Content"] = content

			err = models.UpdatePostById(id, updatePost)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/post/edit/" + id + "/", http.StatusSeeOther)
			return
		}
	}else if r.Method == "GET" {
		post, err := models.FindPostById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if post.UserID != userId {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		collections, err := models.FindAllCollections()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var pageQuery = map[string]interface{}{}
		pageQuery["post"] = post
		pageQuery["collections"] = collections

		pageContent := PageContent{
			PageTitle: "Edit Post",
			PageQuery: pageQuery,
			IsUser: isUser,
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

	post, err := models.FindPostById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if post.UserID != userId {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = models.DeletePostById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/dashboard/", http.StatusSeeOther)
	return
}