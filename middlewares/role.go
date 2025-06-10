package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RequireAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Chỉ admin mới được phép"})
            c.Abort()
            return
        }
        c.Next()
    }
}
