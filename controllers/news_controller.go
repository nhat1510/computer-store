package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"computer-store/models"
	"computer-store/services"
)

// @Summary Lấy danh sách tất cả tin tức
// @Description Trả về danh sách toàn bộ tin tức
// @Tags News
// @Produce json
// @Success 200 {array} models.News
// @Failure 500 {object} map[string]string
// @Router /news [get]
func GetNews(c *gin.Context) {
	news, err := services.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news"})
		return
	}
	c.JSON(http.StatusOK, news)
}

// @Summary Lấy chi tiết tin tức theo ID
// @Description Trả về chi tiết tin tức theo ID
// @Tags News
// @Produce json
// @Param id path int true "ID tin tức"
// @Success 200 {object} models.News
// @Failure 404 {object} map[string]string
// @Router /news/{id} [get]
func GetNewsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	news, err := services.GetNewsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	c.JSON(http.StatusOK, news)
}

// @Summary Tạo tin tức mới
// @Description Tạo một tin tức mới từ form-data (bao gồm ảnh, tiêu đề, nội dung...)
// @Tags News
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Tiêu đề"
// @Param content formData string true "Nội dung"
// @Param highlight formData string false "Nội dung nổi bật"
// @Param tags formData string false "Tags (ngăn cách bằng dấu phẩy)"
// @Param image formData file false "Ảnh đại diện"
// @Success 201 {object} models.News
// @Failure 500 {object} map[string]string
// @Router /news [post]
// @Security BearerAuth
func CreateNews(c *gin.Context) {
	var news models.News

	// ✅ Nhận dữ liệu từ form-data
	news.Title = c.PostForm("title")
	news.Content = c.PostForm("content")
	news.Highlight = c.PostForm("highlight")
	news.Tags = c.PostForm("tags")

	// ✅ Xử lý ảnh nếu có
	file, err := c.FormFile("image")
	if err == nil && file.Filename != "" {
		dst := "./uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		news.Image = "/uploads/" + file.Filename
	}

	created, err := services.CreateNews(&news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news"})
		return
	}
	c.JSON(http.StatusCreated, created)
}


// @Summary Cập nhật tin tức
// @Description Cập nhật tin tức theo ID (form-data)
// @Tags News
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "ID tin tức"
// @Param title formData string false "Tiêu đề"
// @Param content formData string false "Nội dung"
// @Param highlight formData string false "Nội dung nổi bật"
// @Param tags formData string false "Tags"
// @Param image formData file false "Ảnh mới"
// @Success 200 {object} models.News
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news/{id} [put]
// @Security BearerAuth
func UpdateNews(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	existing, err := services.GetNewsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	// ✅ Cập nhật dữ liệu từ form-data
	existing.Title = c.PostForm("title")
	existing.Content = c.PostForm("content")
	existing.Highlight = c.PostForm("highlight")
	existing.Tags = c.PostForm("tags")

	// ✅ Nếu có ảnh mới
	file, err := c.FormFile("image")
	if err == nil && file.Filename != "" {
		dst := "./uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		existing.Image = "/uploads/" + file.Filename
	}

	updated, err := services.UpdateNews(uint(id), existing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// @Summary Xóa tin tức
// @Description Xóa tin tức theo ID
// @Tags News
// @Produce json
// @Param id path int true "ID tin tức"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news/{id} [delete]
// @Security BearerAuth
func DeleteNews(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := services.DeleteNews(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}
