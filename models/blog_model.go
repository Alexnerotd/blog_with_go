package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
