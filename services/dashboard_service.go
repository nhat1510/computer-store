package services

import (
	"computer-store/config"
	"computer-store/models"
	"database/sql"
)

type DashboardData struct {
	TotalOrders  int64             `json:"total_orders"`
	TotalUsers   int64             `json:"total_users"`
	TotalRevenue float64           `json:"total_revenue"`
	TopProducts  []TopProductEntry `json:"top_products"`
}

type TopProductEntry struct {
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	TotalSold int64  `json:"total_sold"`
}

func GetDashboardData() (DashboardData, error) {
	var totalOrders int64
	var totalUsers int64
	var totalRevenue float64

	config.DB.Model(&models.Order{}).Count(&totalOrders)
	config.DB.Model(&models.User{}).Count(&totalUsers)

	var revenue sql.NullFloat64
	config.DB.Table("orders").Select("SUM(total)").Scan(&revenue)
	if revenue.Valid {
		totalRevenue = revenue.Float64
	}

	var topProducts []TopProductEntry
	config.DB.
		Table("order_items").
		Select("products.id as product_id, products.name, SUM(order_items.quantity) as total_sold").
		Joins("JOIN products ON order_items.product_id = products.id").
		Group("products.id, products.name").
		Order("total_sold DESC").
		Limit(5).
		Scan(&topProducts)

	return DashboardData{
		TotalOrders:  totalOrders,
		TotalUsers:   totalUsers,
		TotalRevenue: totalRevenue,
		TopProducts:  topProducts,
	}, nil
}
