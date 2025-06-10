package services

import (
	"computer-store/config"
	"computer-store/models"
	"time"
)

// Lấy user theo ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Cập nhật profile (không avatar)
func UpdateUserProfile(id uint, name, phone, address, city, district, ward, gender string, dob *time.Time, job, bio string) (models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	updates := map[string]interface{}{
		"name":     name,
		"phone":    phone,
		"address":  address,
		"city":     city,
		"district": district,
		"ward":     ward,
		"gender":   gender,
		"job":      job,
		"bio":      bio,
	}

	if dob != nil {
		updates["dob"] = *dob
	}

	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Xóa user theo ID
func DeleteUserByID(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
