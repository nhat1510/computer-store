package controllers

import (
	"computer-store/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Lấy dữ liệu dashboard admin
// @Description Trả về thống kê tổng quan hệ thống như tổng đơn hàng, doanh thu, số lượng người dùng, v.v.
// @Tags Admin
// @Produce json
// @Success 200 {object} services.DashboardData
// @Failure 500 {object} map[string]string
// @Router /admin/dashboard [get]
// @Security BearerAuth
func GetAdminDashboard(c *gin.Context) {
	data, err := services.GetDashboardData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy dữ liệu dashboard"})
		return
	}

	c.JSON(http.StatusOK, data)
}
