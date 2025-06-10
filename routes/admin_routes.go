package routes

import (
    "computer-store/controllers"
    "computer-store/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine) {
    admin := r.Group("/admin")
    admin.Use(middlewares.JWTAuth(), middlewares.RequireAdmin())
    {
        admin.POST("/products", controllers.CreateProduct)
        admin.PUT("/products/:id", controllers.UpdateProduct)
        admin.DELETE("/products/:id", controllers.DeleteProduct)

        admin.GET("/orders", controllers.GetAllOrders)
        admin.GET("/dashboard", controllers.GetAdminDashboard)
    }
}
