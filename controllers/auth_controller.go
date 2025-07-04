package controllers

import (
    "net/http"
    "computer-store/services"
    "computer-store/controllers/types"
    "github.com/gin-gonic/gin"
)

// @Summary Đăng ký tài khoản
// @Description Đăng ký người dùng mới
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body types.RegisterInput true "Thông tin đăng ký"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
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


// @Summary Đăng nhập
// @Description Đăng nhập và nhận JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body types.LoginInput true "Thông tin đăng nhập"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
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

// @Summary Quên mật khẩu
// @Description Gửi mã khôi phục đến email
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body types.ForgotPasswordInput true "Email người dùng"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /forgot-password [post]
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

// @Summary Đặt lại mật khẩu
// @Description Đặt lại mật khẩu bằng mã khôi phục
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body types.ResetPasswordInput true "Email, mã code và mật khẩu mới"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /reset-password [post]
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

// @Summary Đổi mật khẩu
// @Description Người dùng đổi mật khẩu khi đã đăng nhập
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body types.ChangePasswordInput true "Mật khẩu cũ và mới"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /change-password [post]
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

// @Summary Xóa tài khoản
// @Description Xóa tài khoản người dùng hiện tại
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /delete-account [delete]
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

