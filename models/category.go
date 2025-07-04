package models

import (
	"time"
)

type Category struct {
	ID        uint       `json:"id" example:"1"`
	Name      string     `json:"name" example:"Laptop"`
	Image     string     `json:"image" example:"laptop.png"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// ✅ Quan hệ 1-n với Product
	Products []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

// CategoryResponse chỉ dùng cho Swagger docs (nếu không cần Products)
type CategoryResponse struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"Điện thoại"`
	Image string `json:"image" example:"phone.jpg"`
}
