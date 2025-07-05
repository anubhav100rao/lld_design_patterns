package controller

import (
	"net/http"

	"github.com/anubhav100rao/lld_design_patterns/mvc/model"
	"github.com/anubhav100rao/lld_design_patterns/mvc/view"
)

// ListArticles handles GET /articles
func ListArticles(w http.ResponseWriter, r *http.Request) {
	articles := model.GetAll()
	view.Render(w, "list", articles)
}

// CreateArticle handles POST /articles
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/articles", http.StatusSeeOther)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	_, err := model.Create(title, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}
