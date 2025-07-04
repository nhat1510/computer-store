package models

import (
	"time"
)

type Review struct {
	ID        uint       `json:"id" gorm:"primaryKey" example:"1"`
	ProductID uint       `json:"product_id" example:"101"`
	UserID    uint       `json:"user_id" example:"5"`
	Rating    int        `json:"rating" example:"4"`
	Comment   string     `json:"comment" example:"Sản phẩm rất tốt"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}
