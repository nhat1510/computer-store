package controllers

import (
	"computer-store/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAdminDashboard(c *gin.Context) {
	data, err := services.GetDashboardData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy dữ liệu dashboard"})
		return
	}

	c.JSON(http.StatusOK, data)
}
