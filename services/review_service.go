package services

import (
    "computer-store/config"
    "computer-store/models"
    "errors"
)

func GetReviewsByProduct(productID uint) ([]models.Review, error) {
    var reviews []models.Review
    err := config.DB.Where("product_id = ?", productID).Find(&reviews).Error
    return reviews, err
}

func GetReviewByID(id uint) (models.Review, error) {
    var review models.Review
    err := config.DB.First(&review, id).Error
    if err != nil {
        return review, errors.New("không tìm thấy đánh giá")
    }
    return review, nil
}

func CreateReview(review *models.Review) error {
    return config.DB.Create(review).Error
}

func UpdateReview(id uint, input *models.Review) error {
    var review models.Review
    if err := config.DB.First(&review, id).Error; err != nil {
        return err
    }

    review.Rating = input.Rating
    review.Comment = input.Comment

    return config.DB.Save(&review).Error
}

func DeleteReview(id uint) error {
    return config.DB.Delete(&models.Review{}, id).Error
}
