package routes

import (
    "computer-store/controllers"
    "computer-store/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupCartRoutes(r *gin.Engine) {
    cart := r.Group("/cart")
    cart.Use(middlewares.JWTAuth())
    {
        cart.POST("", controllers.AddToCart)
        cart.GET("", controllers.GetCart)
        cart.PUT("/:product_id", controllers.UpdateCartItem)
        cart.DELETE("/:product_id", controllers.RemoveFromCart)
    }
}
