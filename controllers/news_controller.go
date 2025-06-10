package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"computer-store/models"
	"computer-store/services"
)

func GetNews(c *gin.Context) {
	news, err := services.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news"})
		return
	}
	c.JSON(http.StatusOK, news)
}

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

func DeleteNews(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := services.DeleteNews(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}
