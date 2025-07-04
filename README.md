#  Computer Store

**Computer Store** là hệ thống backend RESTful API hỗ trợ quản lý và bán hàng linh kiện máy tính.
---
##  Tính năng chính

###  Người dùng:
- Đăng ký / Đăng nhập (JWT)
- Xem danh sách & chi tiết sản phẩm
- Tìm kiếm sản phẩm theo từ khóa
- Quản lý giỏ hàng
- Tạo đơn hàng
- Đánh giá sản phẩm
- Xem đơn hàng cá nhân
- thêm thông tin của khách hàng

###  Admin:
- Quản lý sản phẩm (thêm, sửa, xoá)
- Quản lý danh mục sản phẩm
- Quản lý người dùng & đơn hàng
- Dashboard thống kê (sẽ bổ sung sau)


##  Công nghệ sử dụng

| Thành phần         | Công nghệ                    |
|-------------------|-------------------------------|
| Ngôn ngữ           | Golang 1.22+                  |
| Web framework      | Gin-Gonic                     |
| ORM                | GORM                          |
| Cơ sở dữ liệu      | PostgreSQL                    |                   
| Authentication     | JWT (golang-jwt)              |
| Email              | Gomail (đang trong quá trình xử lý)|
| Cấu hình runtime   | `.env`, `.air.toml`           |
| Tổ chức code       | theo mô hình controller-service-model|

---

##  Hướng dẫn chạy dự án
- go mod tidy
- go run main.go
- air

### 1. Cài đặt yêu cầu
- Golang >= 1.22
- PostgreSQL

### 2. Cấu hình `.env`
Tạo file `.env` (nếu chưa có) 
DB_USER=postgres
DB_PASSWORD=your_password
DB_HOST=localhost
DB_NAME=computer_store
DB_PORT=5432

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASS=your_email_password

JWT_SECRET=mysecretkey

Tạo thư mục upload/ (nếu muốn đăng ảnh)

