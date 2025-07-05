package model

import "errors"

// Article represents a simple data entity.
type Article struct {
	ID      int
	Title   string
	Content string
}

var (
	articles = []Article{}
	nextID   = 1
)

// GetAll returns all articles.
func GetAll() []Article {
	return articles
}

// Create adds a new article.
func Create(title, content string) (Article, error) {
	if title == "" {
		return Article{}, errors.New("title cannot be empty")
	}
	art := Article{
		ID:      nextID,
		Title:   title,
		Content: content,
	}
	nextID++
	articles = append(articles, art)
	return art, nil
}
