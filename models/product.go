package models

import "time"

type Product struct {
	ID          uint       `json:"id" example:"1"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Stock       int        `json:"stock"`
	Image       string     `json:"image"`
	ImageURL    string     `json:"image_url" gorm:"-"` // không lưu DB, tự tạo khi lấy dữ liệu
	CategoryID  uint       `json:"category_id"`
	Category    Category   `json:"category"` // sẽ dùng struct sửa như trên
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

