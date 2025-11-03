package models

import "gorm.io/gorm"

type Sticker struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price" gorm:"not null"`
}
