package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Image    string   `json:"image"`
	Products []Product      `json:"products" gorm:"foreignKey:CategoryID"`
}
