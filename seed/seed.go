package seed

import (
	"computer-store/config"
	"computer-store/models"
	"computer-store/utils"
	"fmt"
)

func SeedAdminUser() {
	// Kiểm tra xem đã có admin chưa
	var admin models.User
	if err := config.DB.Where("email = ?", "admin@gmail.com").First(&admin).Error; err == nil {
		fmt.Println("ngon Admin đã tồn tại, bỏ qua seed")
		return
	}

	// Mã hóa mật khẩu
	hashedPassword, _ := utils.HashPassword("123456")

	admin = models.User{
		Name:     "Super Admin",
		Email:    "admin@gmail.com",
		Password: hashedPassword,
		Role:     "admin",
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		fmt.Println("❌ Lỗi khi tạo admin:", err)
		return
	}

	fmt.Println("✅ Đã seed admin thành công (email: admin@gmail.com, pass: 123456)")
}
