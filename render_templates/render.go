package render_templates

import (
	"golang_side_project_crud_website/controller"
	"html/template"
	"net/http"
	"path"
)


func ReturnRenderTemplate(w http.ResponseWriter,
	templatePath string,
	pageContent *controller.PageContent) {

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