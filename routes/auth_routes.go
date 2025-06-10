package routes

import (
    "computer-store/controllers"
    "github.com/gin-gonic/gin"
    "computer-store/middlewares"

)

func SetupAuthRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.PUT("/reset-password", middlewares.JWTAuth(), controllers.ChangePassword)
    r.DELETE("/delete-account", middlewares.JWTAuth(), controllers.DeleteAccount)

   
}
