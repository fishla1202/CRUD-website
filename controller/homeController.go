package controller

import (
	"html/template"
	"net/http"
	"path"
)

type PageContent struct {
	PageTitle string
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	layout := path.Join("templates", "layout.html")
	index := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(layout, index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageContent := PageContent{PageTitle: "FishLa Blog"}

	err = tmpl.Execute(w, &pageContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
