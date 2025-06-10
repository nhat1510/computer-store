package routes

import (
	"computer-store/controllers"
	"computer-store/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupNewsRoutes(r *gin.Engine) {
	api := r.Group("/api")
	news := api.Group("/news")

	// Public
	news.GET("", controllers.GetNews)
	news.GET("/:id", controllers.GetNewsByID)

	// Admin-only
	newsAdmin := news.Group("")
	newsAdmin.Use(middlewares.JWTAuth(), middlewares.RequireAdmin())
	{
		newsAdmin.POST("", controllers.CreateNews)
		newsAdmin.PUT("/:id", controllers.UpdateNews)
		newsAdmin.DELETE("/:id", controllers.DeleteNews)
	}
}
