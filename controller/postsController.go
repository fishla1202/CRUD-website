package controller

import (
	"CRUD/models/posts"
	"log"
	"net/http"
)


func PostIndex(w http.ResponseWriter, r *http.Request) {

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
