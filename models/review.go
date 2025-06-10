package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}
