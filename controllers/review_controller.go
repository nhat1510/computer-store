package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "computer-store/config"
    "computer-store/models"
    "computer-store/services"
)

func GetReviewsByProductID(c *gin.Context) {
    idParam := c.Param("product_id")
    productID, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID sản phẩm không hợp lệ"})
        return
    }

    var reviews []models.Review
    
    if err := config.DB.Preload("User").Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy vấn đánh giá"})
        return
    }

    c.JSON(http.StatusOK, reviews)
}

// POST /reviews
func CreateReview(c *gin.Context) {
    var input models.Review
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
        return
    }

    // ✅ Lấy user_id từ JWT middleware (gán sẵn vào context)
    userID := c.GetUint("user_id")
    input.UserID = userID

    if err := services.CreateReview(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo đánh giá"})
        return
    }

    c.JSON(http.StatusCreated, input)
}

// PUT /reviews/:id
func UpdateReview(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var input models.Review
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
        return
    }

    if err := services.UpdateReview(uint(id), &input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đã cập nhật thành công"})
}

// DELETE /reviews/:id
func DeleteReview(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := services.DeleteReview(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đã xóa thành công"})
}
