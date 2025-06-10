package services

import (
	"computer-store/config"
	"computer-store/models"
	"errors"
)

// Tạo sản phẩm từ form (upload ảnh)
func CreateProductFromForm(input models.Product) (models.Product, error) {
	if input.Name == "" || input.Image == "" {
		return models.Product{}, errors.New("Tên và ảnh sản phẩm là bắt buộc")
	}
	if err := config.DB.Create(&input).Error; err != nil {
		return models.Product{}, err
	}
	return input, nil
}

// Lấy tất cả sản phẩm
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Lấy sản phẩm theo ID
func GetProductByID(id string) (models.Product, error) {
	var product models.Product
	if err := config.DB.Preload("Category").First(&product, "id = ?", id).Error; err != nil {
		return models.Product{}, errors.New("Không tìm thấy sản phẩm")
	}
	return product, nil
}

// Xóa sản phẩm
func DeleteProduct(id string) error {
	if err := config.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		return errors.New("Xóa sản phẩm thất bại")
	}
	return nil
}
