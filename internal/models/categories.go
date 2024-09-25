package models

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,lte=255"`
	CreatedAt time.Time `json:"created_at"`
}
