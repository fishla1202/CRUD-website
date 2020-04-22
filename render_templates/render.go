package render_templates

import (
	"html/template"
	"net/http"
	"path"
)

func ReturnRenderTemplate(w http.ResponseWriter,
	templatePath string,
	pageContent interface{}) {
	layout := path.Join("templates/layout", "layout.html")
	sidebar := path.Join("templates/layout", "sidebar.html")
	postList := path.Join("templates/layout", "post_list.html")
	tmpl, err := template.ParseFiles(layout, sidebar, postList, templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, pageContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
