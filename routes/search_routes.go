package routes

import (
	"computer-store/controllers"
	"github.com/gin-gonic/gin"
)

func SetupSearchRoutes(r *gin.Engine) {
	r.GET("/search/products", controllers.SearchProductsHandler)
}
