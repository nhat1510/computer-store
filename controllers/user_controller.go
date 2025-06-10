package controllers

import (
	"computer-store/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Lấy thông tin tài khoản hiện tại
func GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Cập nhật thông tin tài khoản (KHÔNG avatar)
func UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	// BIND THÀNH CHUỖI TRƯỚC
	var input struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		City     string `json:"city"`
		District string `json:"district"`
		Ward     string `json:"ward"`
		Gender   string `json:"gender"`
		Job      string `json:"job"`
		Bio      string `json:"bio"`
		Dob      string `json:"dob"` // string để parse
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu gửi lên không hợp lệ"})
		return
	}

	// Parse dob nếu có
	var dobPtr *time.Time
	if input.Dob != "" {
		dob, err := time.Parse("2006-01-02", input.Dob)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ngày sinh không đúng định dạng yyyy-mm-dd"})
			return
		}
		dobPtr = &dob
	}

	user, err := services.UpdateUserProfile(
	userID,
	input.Name,
	input.Phone,
	input.Address,
	input.City,
	input.District,
	input.Ward,
	input.Gender,
	dobPtr,
	input.Job,
	input.Bio,
)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cập nhật thất bại"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Xóa tài khoản
func DeleteProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	if err := services.DeleteUserByID(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa tài khoản"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tài khoản đã được xóa"})
}
