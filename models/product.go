package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
    ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Image       string   `json:"image"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}
