package routes

import (
    "time"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	r.Static("/uploads", "./uploads")

    SetupAuthRoutes(r)
    SetupProductRoutes(r)
    SetupOrderRoutes(r)
    SetupCartRoutes(r)
    SetupReviewRoutes(r)
    SetupAdminRoutes(r)
    SetupCategoryRoutes(r)
    SetupSearchRoutes(r)
    SetupNewsRoutes(r)
    SetupUserRoutes(r)

    return r
}
