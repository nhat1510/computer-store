package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "computer-store/config"
    "computer-store/models"
    "computer-store/services"
)

// @Summary Lấy đánh giá theo sản phẩm
// @Description Trả về danh sách đánh giá của một sản phẩm cụ thể (bao gồm thông tin người dùng)
// @Tags Reviews
// @Produce json
// @Param product_id path int true "ID sản phẩm"
// @Success 200 {array} models.Review
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/product/{product_id} [get]
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

// @Summary Tạo đánh giá mới
// @Description Người dùng tạo đánh giá cho một sản phẩm
// @Tags Reviews
// @Accept json
// @Produce json
// @Param input body models.Review true "Thông tin đánh giá"
// @Success 201 {object} models.Review
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews [post]
// @Security BearerAuth
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

// @Summary Cập nhật đánh giá
// @Description Cập nhật nội dung hoặc điểm đánh giá theo ID
// @Tags Reviews
// @Accept json
// @Produce json
// @Param id path int true "ID đánh giá"
// @Param input body models.Review true "Thông tin đánh giá cập nhật"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [put]
// @Security BearerAuth
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

// @Summary Xoá đánh giá
// @Description Xoá một đánh giá theo ID
// @Tags Reviews
// @Produce json
// @Param id path int true "ID đánh giá"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [delete]
// @Security BearerAuth
func DeleteReview(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := services.DeleteReview(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đã xóa thành công"})
}
