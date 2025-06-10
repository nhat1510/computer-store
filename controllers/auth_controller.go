package controllers

import (
    "net/http"
    "computer-store/services"
    "computer-store/controllers/types"
    "github.com/gin-gonic/gin"
)


func Register(c *gin.Context) {
    var input types.RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.RegisterUser(input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
}


func Login(c *gin.Context) {
    var input types.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, user, err := services.LoginUser(input)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user": gin.H{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
            "role":  user.Role,
        },
    })
}


func ForgotPassword(c *gin.Context) {
    var input types.ForgotPasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.SendResetCode(input.Email); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Mã khôi phục đã được gửi đến email"})
}


func ResetPassword(c *gin.Context) {
    var input types.ResetPasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.VerifyAndResetPassword(input.Email, input.Code, input.NewPassword); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Mật khẩu đã được đặt lại thành công"})
}


func ChangePassword(c *gin.Context) {
    var input types.ChangePasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userIDVal, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Không xác định được người dùng"})
        return
    }

    userID, ok := userIDVal.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "ID người dùng không hợp lệ"})
        return
    }

    if err := services.ChangeUserPassword(userID, input.OldPassword, input.NewPassword); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đổi mật khẩu thành công"})
}

func DeleteAccount(c *gin.Context) {
    userIDVal, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Không xác định được người dùng"})
        return
    }

    userID, ok := userIDVal.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "ID người dùng không hợp lệ"})
        return
    }

    if err := services.DeleteUserByID(userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Tài khoản đã được xóa"})
}

