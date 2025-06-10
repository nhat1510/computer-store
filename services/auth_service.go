package services

import (
    "errors"
    "fmt"
    "math/rand"
    "net/smtp"
    "computer-store/config"
    "computer-store/models"
    "computer-store/utils"
    "computer-store/controllers/types"
)

var resetCodes = make(map[string]string) // lưu tạm mã xác thực theo email

func LoginUser(input types.LoginInput) (string, models.User, error) {
    var user models.User

    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        return "", user, errors.New("Email không tồn tại")
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        return "", user, errors.New("Mật khẩu không đúng")
    }

    token, err := utils.GenerateJWT(user.ID, user.Role)
    if err != nil {
        return "", user, errors.New("Không thể tạo token")
    }

    return token, user, nil
}

func RegisterUser(input types.RegisterInput) error {
    var existing models.User
    if err := config.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
        return errors.New("Email đã được sử dụng")
    }

    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        return errors.New("Không thể mã hóa mật khẩu")
    }

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: hashedPassword,
        Role:     "user",
    }

    if err := config.DB.Create(&user).Error; err != nil {
        return errors.New("Không thể tạo người dùng")
    }

    return nil
}

// ✅ Gửi mã khôi phục qua email
func SendResetCode(email string) error {
    var user models.User
    if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return errors.New("Email không tồn tại")
    }

    code := fmt.Sprintf("%06d", rand.Intn(1000000))
    resetCodes[email] = code

    subject := "Mã khôi phục mật khẩu"
    body := fmt.Sprintf("Xin chào %s,\n\nMã khôi phục mật khẩu của bạn là: %s\n\nVui lòng không chia sẻ mã này với người khác.", user.Name, code)

    return sendEmail(email, subject, body)
}

// ✅ Đặt lại mật khẩu bằng mã đã gửi
func VerifyAndResetPassword(email, code, newPassword string) error {
    if resetCodes[email] != code {
        return errors.New("Mã xác thực không đúng hoặc đã hết hạn")
    }

    hashedPassword, err := utils.HashPassword(newPassword)
    if err != nil {
        return errors.New("Không thể mã hóa mật khẩu mới")
    }

    if err := config.DB.Model(&models.User{}).
        Where("email = ?", email).
        Update("password", hashedPassword).Error; err != nil {
        return errors.New("Không thể cập nhật mật khẩu")
    }

    delete(resetCodes, email)
    return nil
}

// ✅ Đổi mật khẩu khi đã đăng nhập
func ChangeUserPassword(userID uint, oldPassword, newPassword string) error {
    var user models.User
    if err := config.DB.First(&user, userID).Error; err != nil {
        return errors.New("Không tìm thấy người dùng")
    }

    if !utils.CheckPasswordHash(oldPassword, user.Password) {
        return errors.New("Mật khẩu cũ không đúng")
    }

    hashed, err := utils.HashPassword(newPassword)
    if err != nil {
        return errors.New("Không thể mã hóa mật khẩu mới")
    }

    return config.DB.Model(&user).Update("password", hashed).Error
}

// ✅ Gửi mail qua SMTP
func sendEmail(to, subject, body string) error {
    from := config.SMTPUser
    pass := config.SMTPPass
    host := config.SMTPHost
    port := config.SMTPPort

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" + body

    auth := smtp.PlainAuth("", from, pass, host)
    return smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
}

// ✅ Xóa tài khoản người dùng theo ID
func  DeleteAccount(userID uint) error {
    // Xóa đánh giá liên quan trước
    if err := config.DB.Where("user_id = ?", userID).Delete(&models.Review{}).Error; err != nil {
        return errors.New("Không thể xóa review của người dùng")
    }

    // Xóa user
    if err := config.DB.Delete(&models.User{}, userID).Error; err != nil {
        return errors.New("Không thể xóa tài khoản")
    }

    return nil
}

