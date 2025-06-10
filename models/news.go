package models

import "time"

type News struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Image     string    `json:"image"` // ảnh đại diện
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Highlight string    `json:"highlight"` // ✅ cần có
	Tags      string    `json:"tags"`  
}

func (News) TableName() string {
	return "news"
}
