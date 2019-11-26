package controller

import (
	"fmt"
	"log"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil { log.Fatal(err)}
		fmt.Println(r.Form["content"])
	}else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
