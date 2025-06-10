package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Total  float64 `json:"total"`
	Status string  `json:"status"`

	// Quan hệ với OrderItem, khi xóa đơn hàng thì xóa luôn item
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`

	// Quan hệ với Product để lấy ảnh hoặc tên
	Product Product `json:"-" gorm:"foreignKey:ProductID"`

	// Field phụ trợ để trả ảnh ra JSON, không lưu DB
	ProductImage string `json:"product_image" gorm:"-"`
	ProductName  string `json:"product_name" gorm:"-"`
}
