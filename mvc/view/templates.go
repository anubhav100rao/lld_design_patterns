package view

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(
	template.ParseFiles(
		filepath.Join("/Users/anubhav100rao/development/lld_design_patterns/mvc/view/templates", "list.html"),
	),
)

// Render executes a named template with data.
func Render(w http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
