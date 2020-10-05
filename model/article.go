package model

import (
	"html/template"
	"time"
)

//ArticleAuthor model
type ArticleAuthor struct {
	ID     int64   `json:"id"`
	Name   *string `json:"name"`
	Gender *string `json:"gender"`
	Dob    *string `json:"dob"`
	Photo  *string `json:"photo"`
}

//Article model
type Article struct {
	ID        int64         `json:"id"`
	Author    ArticleAuthor `json:"author"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Summary   string        `json:"summary"`
	Images    string        `json:"images"`
	Tags      string        `json:"tags"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

// ArticleMutation used
type ArticleMutation struct {
	ID        *int64     `json:"id,omitempty"`
	AuthorID  *int64     `json:"authorId,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Content   *string    `json:"content,omitempty"`
	Summary   *string    `json:"summary,omitempty"`
	Images    *string    `json:"images,omitempty"`
	Tags      *string    `json:"tags,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// ArticleQuery used
type ArticleQuery struct {
	Page         *int64  `form:"page"`
	ItemsPerPage *int64  `form:"itemsPerPage"`
	SortBy       *string `form:"sortBy"`
	SortOrder    *string `form:"sortOrder"`
}

// HTMLContent func
func (a *Article) HTMLContent() template.HTML {
	return template.HTML(a.Content)
}

// FormattedCreateDate func
func (a *Article) FormattedCreateDate() string {
	return a.CreatedAt.Local().Format("Monday, 02 Jan 15:04")
}
