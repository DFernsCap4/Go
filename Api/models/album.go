package models

import(
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model // adds fields, ID, createdAt, UpdatedAt, DeletedAt
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}
