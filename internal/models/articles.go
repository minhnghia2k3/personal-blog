package models

import (
	"html/template"
	"time"
)

type Article struct {
	ID        int           `json:"id"`
	Title     string        `json:"title" validate:"required,lte=255"`
	Content   template.HTML `json:"content" validate:"required"`
	MinRead   int           `json:"min_read" validate:"required,min=1"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}
