package services

import (
    "errors"
    "computer-store/config"
    "computer-store/models"
)

func CreateCategory(input models.Category) (models.Category, error) {
    if err := config.DB.Create(&input).Error; err != nil {
        return models.Category{}, errors.New("Không thể tạo danh mục")
    }
    return input, nil
}

func GetAllCategories() ([]models.Category, error) {
    var categories []models.Category
    if err := config.DB.Preload("Products").Find(&categories).Error; err != nil {
        return nil, errors.New("Lỗi lấy danh sách danh mục")
    }
    return categories, nil
}

func UpdateCategory(id string, updated models.Category) (*models.Category, error) {
	var category models.Category

	// Kiểm tra danh mục tồn tại
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, errors.New("Danh mục không tồn tại")
	}

	// Cập nhật tên
	category.Name = updated.Name

	// Nếu có ảnh mới thì cập nhật
	if updated.Image != "" {
		category.Image = updated.Image
	}

	if err := config.DB.Save(&category).Error; err != nil {
		return nil, errors.New("Không thể cập nhật danh mục")
	}

	return &category, nil
}


func DeleteCategory(id string) error {
	var category models.Category

	// Kiểm tra danh mục tồn tại
	if err := config.DB.First(&category, id).Error; err != nil {
		return errors.New("Danh mục không tồn tại")
	}

	// Xóa
	if err := config.DB.Delete(&category).Error; err != nil {
		return errors.New("Không thể xoá danh mục")
	}

	return nil
}

func GetProductsByCategorySlug(slug string) ([]models.Product, error) {
	var category models.Category
	err := config.DB.Where("slug = ?", slug).First(&category).Error
	if err != nil {
		return nil, errors.New("Không tìm thấy danh mục")
	}

	var products []models.Product
	err = config.DB.Where("category_id = ?", category.ID).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}