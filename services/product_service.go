package services

import (
	"computer-store/config"
	"computer-store/models"
	"encoding/json"
	"errors"
	"time"
)

func CreateProductFromForm(input models.Product) (models.Product, error) {
	if input.Name == "" || input.Image == "" {
		return models.Product{}, errors.New("Tên và ảnh sản phẩm là bắt buộc")
	}
	if err := config.DB.Create(&input).Error; err != nil {
		return models.Product{}, err
	}

	config.RedisClient.Del(config.Ctx, "products:all")
	return input, nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	cached, err := config.RedisClient.Get(config.Ctx, "products:all").Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cached), &products); err == nil {
			return products, nil
		}
	}

	if err := config.DB.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}

	data, _ := json.Marshal(products)
	config.RedisClient.Set(config.Ctx, "products:all", data, time.Minute*10)

	return products, nil
}

func GetProductByID(id string) (models.Product, error) {
	var product models.Product
	if err := config.DB.Preload("Category").First(&product, "id = ?", id).Error; err != nil {
		return models.Product{}, errors.New("Không tìm thấy sản phẩm")
	}
	return product, nil
}

func DeleteProduct(id string) error {
	if err := config.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		return errors.New("Xóa sản phẩm thất bại")
	}
	config.RedisClient.Del(config.Ctx, "products:all")
	return nil
}
