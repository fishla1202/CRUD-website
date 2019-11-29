package render_templates

import (
	"html/template"
	"net/http"
	"path"
)


func ReturnRenderTemplate(w http.ResponseWriter,
	templatePath string,
	pageContent interface{}) {

	layout := path.Join("templates", "layout.html")
	tmpl, err := template.ParseFiles(layout, templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, pageContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}