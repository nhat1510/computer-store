package controllers

import (
	"computer-store/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchProductsHandler(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Từ khoá không được để trống"})
		return
	}

	products, err := services.SearchProductsSimple(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi tìm kiếm sản phẩm"})
		return
	}

	c.JSON(http.StatusOK, products)
}
