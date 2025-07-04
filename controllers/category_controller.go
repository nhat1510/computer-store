package controllers

import (
    "net/http"
    "computer-store/models"
    "computer-store/services"
    "github.com/gin-gonic/gin"
)

// @Summary Tạo danh mục mới
// @Description Tạo mới danh mục với tên và ảnh đại diện
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Tên danh mục"
// @Param image formData file true "Ảnh đại diện danh mục"
// @Success 200 {array} models.CategoryResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
// @Security BearerAuth
func CreateCategory(c *gin.Context) {
    name := c.PostForm("name")

    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Vui lòng chọn ảnh"})
        return
    }

    // CHỈ LƯU TÊN FILE (KHÔNG CÓ "uploads/")
    imagePath := file.Filename

    if err := c.SaveUploadedFile(file, "uploads/"+imagePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu ảnh"})
        return
    }

    category := models.Category{
        Name:  name,
        Image: imagePath, // chỉ là tên file
    }

    created, err := services.CreateCategory(category)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, created)
}

// @Summary Cập nhật danh mục
// @Description Cập nhật thông tin danh mục theo ID, có thể đổi tên và/hoặc ảnh
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID danh mục"
// @Param name formData string false "Tên mới"
// @Param image formData file false "Ảnh mới (tuỳ chọn)"
// @Success 200 {array} models.CategoryResponse
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [put]
// @Security BearerAuth
func UpdateCategory(c *gin.Context) {
    id := c.Param("id")
    name := c.PostForm("name")
    imagePath := ""

    file, err := c.FormFile("image")
    if err == nil {
        imagePath = file.Filename
        if err := c.SaveUploadedFile(file, "uploads/"+imagePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu ảnh"})
            return
        }
    }

    category := models.Category{
        Name:  name,
        Image: imagePath, // nếu rỗng thì giữ nguyên ảnh cũ trong service
    }

    updated, err := services.UpdateCategory(id, category)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updated)
}

// @Summary Lấy tất cả danh mục
// @Description Trả về danh sách tất cả danh mục hiện có
// @Tags Categories
// @Produce json
 // @Success 200 {array} models.CategoryResponse
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
    categories, err := services.GetAllCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, categories)
}

// @Summary Xoá danh mục
// @Description Xoá danh mục theo ID
// @Tags Categories
// @Produce json
// @Param id path string true "ID danh mục"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
// @Security BearerAuth
func DeleteCategory(c *gin.Context) {
    id := c.Param("id")

    if err := services.DeleteCategory(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Xoá danh mục thành công"})
}
