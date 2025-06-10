package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Email     string     `gorm:"type:varchar(100);unique" json:"email"`
	Password  string     `gorm:"type:varchar(255)" json:"-"`
	Phone     string     `gorm:"type:varchar(20)" json:"phone,omitempty"`
	Address   string     `gorm:"type:text" json:"address,omitempty"`
	City      string     `gorm:"type:varchar(100)" json:"city,omitempty"`
	District  string     `gorm:"type:varchar(100)" json:"district,omitempty"`
	Ward      string     `gorm:"type:varchar(100)" json:"ward,omitempty"`
	Avatar    string     `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	Gender    string     `gorm:"type:varchar(10)" json:"gender,omitempty"` 
	Dob       *time.Time `json:"dob,omitempty"`  
	Job       string     `gorm:"type:varchar(100)" json:"job,omitempty"`
	Bio       string     `gorm:"type:text" json:"bio,omitempty"`
	Role      string     `gorm:"default:user" json:"role"`
	Status    string     `gorm:"type:varchar(20);default:'active'" json:"status"` 
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
