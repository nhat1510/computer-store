package routes

import (
	"computer-store/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.GET("", controllers.GetAllProducts)
		products.GET("/:id", controllers.GetProductByID)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}
}
