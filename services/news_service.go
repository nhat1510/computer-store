package services

import (
	"errors"
	"computer-store/config"
	"computer-store/models"
)

func GetAllNews() ([]models.News, error) {
	var news []models.News
	if err := config.DB.Order("created_at desc").Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func GetNewsByID(id uint) (*models.News, error) {
	var news models.News
	if err := config.DB.First(&news, id).Error; err != nil {
		return nil, errors.New("news not found")
	}
	return &news, nil
}

func CreateNews(input *models.News) (*models.News, error) {
	if err := config.DB.Create(input).Error; err != nil {
		return nil, err
	}
	return input, nil
}

func UpdateNews(id uint, input *models.News) (*models.News, error) {
	news, err := GetNewsByID(id)
	if err != nil {
		return nil, err
	}
	if err := config.DB.Model(&news).Updates(input).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func DeleteNews(id uint) error {
	if err := config.DB.Delete(&models.News{}, id).Error; err != nil {
		return err
	}
	return nil
}
