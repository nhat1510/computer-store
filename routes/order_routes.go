package routes

import (
    "computer-store/controllers"
    "computer-store/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupOrderRoutes(r *gin.Engine) {
    user := r.Group("/orders")
    user.Use(middlewares.JWTAuth())
    {
        user.POST("", controllers.CreateOrder)
        user.GET("", controllers.GetUserOrders)
        user.POST("/from-cart", controllers.CreateOrderFromCart)
    }
}
