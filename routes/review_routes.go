package routes

import (
    "computer-store/controllers"
    "computer-store/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupReviewRoutes(r *gin.Engine) {
    review := r.Group("/reviews")
    {
        review.GET("/by-product/:product_id", controllers.GetReviewsByProductID)

        review.Use(middlewares.JWTAuth())
        {
            review.POST("", controllers.CreateReview)
            review.PUT("/:id", controllers.UpdateReview)
            review.DELETE("/:id", controllers.DeleteReview)
        }
    }
}

