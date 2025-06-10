package services

import (
	"computer-store/config"
	"computer-store/models"
)

func SearchProductsSimple(keyword string) ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.
		Preload("Category").
		Where("name ILIKE ?", "%"+keyword+"%").
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
