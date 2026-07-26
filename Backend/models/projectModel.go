package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string `gorm:"size:100" json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"omitempty"`
}
