package main

import (
	"net/http"

	"github.com/anubhav100rao/lld_design_patterns/mvc/controller"
)

func main() {
	http.HandleFunc("/articles/create", controller.CreateArticle) // or handle POST on /articles
	http.HandleFunc("/articles", controller.ListArticles)

	http.ListenAndServe(":8080", nil)
}
