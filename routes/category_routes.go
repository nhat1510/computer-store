package routes

import (
	"github.com/gin-gonic/gin"
	"computer-store/controllers"
	"computer-store/middlewares"
	
)

func SetupCategoryRoutes(r *gin.Engine) {
	public := r.Group("/categories")
	{
		public.GET("", controllers.GetAllCategories) // Public API
		
	}

	admin := r.Group("/admin")
	admin.Use(middlewares.JWTAuth(), middlewares.RequireAdmin())
	{
		admin.POST("/categories", controllers.CreateCategory)
		admin.PUT("/categories/:id", controllers.UpdateCategory)
		admin.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
