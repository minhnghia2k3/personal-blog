package models

import (
	"html/template"
	"time"
)

type Article struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Content   template.HTML `json:"content"`
	MinRead   int           `json:"min_read"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}
