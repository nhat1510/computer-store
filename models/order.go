package models

import (
	"time"
)

// Order đại diện cho đơn hàng
type Order struct {
	ID        uint         `json:"id" gorm:"primaryKey" example:"1"`
	UserID    uint         `json:"user_id" example:"2"`
	Total     float64      `json:"total" example:"1500000"`
	Status    string       `json:"status" example:"Đang xử lý"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt *time.Time   `json:"deleted_at,omitempty"`

	// ➕ Danh sách sản phẩm trong đơn hàng
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

// OrderItem là chi tiết đơn hàng
type OrderItem struct {
	ID        uint       `json:"id" gorm:"primaryKey" example:"1"`
	OrderID   uint       `json:"order_id" example:"10"`
	ProductID uint       `json:"product_id" example:"5"`
	Quantity  int        `json:"quantity" example:"2"`
	Price     float64    `json:"price" example:"750000"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// ➖ Không trả về toàn bộ Product
	Product Product `json:"-" gorm:"foreignKey:ProductID"`

	// ➕ Trả riêng tên & ảnh sản phẩm
	ProductImage string `json:"product_image" gorm:"-" example:"laptop.jpg"`
	ProductName  string `json:"product_name" gorm:"-" example:"Laptop Acer Gaming"`
}
