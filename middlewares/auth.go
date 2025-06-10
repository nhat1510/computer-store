package middlewares

import (
    "computer-store/utils"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "net/http"
    "strings"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Thiếu token"})
            c.Abort()
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        token, err := jwt.ParseWithClaims(tokenStr, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte("your-secret-key"), nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ", "detail": err.Error()})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(*utils.Claims)
        if !ok || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}
