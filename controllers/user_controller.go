package controllers

import (
	"computer-store/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary Lấy thông tin tài khoản
// @Description Lấy thông tin của người dùng đang đăng nhập
// @Tags Profile
// @Produce json
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /profile [get]
// @Security BearerAuth
func GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy người dùng"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Cập nhật thông tin tài khoản
// @Description Cập nhật thông tin tài khoản người dùng (không bao gồm avatar)
// @Tags Profile
// @Accept json
// @Produce json
// @Param input body object true "Thông tin người dùng"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profile [put]
// @Security BearerAuth
func UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

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

// @Summary Xoá tài khoản người dùng
// @Description Xoá tài khoản hiện tại và toàn bộ dữ liệu liên quan
// @Tags Profile
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profile [delete]
// @Security BearerAuth
func DeleteProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	if err := services.DeleteUserByID(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa tài khoản"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tài khoản đã được xóa"})
}
