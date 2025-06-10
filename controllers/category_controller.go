package controllers

import (
    "net/http"
    "computer-store/models"
    "computer-store/services"
    "github.com/gin-gonic/gin"
)

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

func GetAllCategories(c *gin.Context) {
    categories, err := services.GetAllCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, categories)
}

func DeleteCategory(c *gin.Context) {
    id := c.Param("id")

    if err := services.DeleteCategory(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Xoá danh mục thành công"})
}
