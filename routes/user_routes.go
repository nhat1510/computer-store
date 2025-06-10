package routes

import (
	"computer-store/controllers"
	"computer-store/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	

	profile := r.Group("/profile")
	profile.Use(middlewares.JWTAuth())
	{
		profile.GET("", controllers.GetProfile)
		profile.PUT("", controllers.UpdateProfile)     
		profile.DELETE("", controllers.DeleteProfile)
	}
}