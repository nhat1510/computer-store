package controllers

import (
    "computer-store/config"
    "computer-store/models"
    "computer-store/services"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// @Summary Tạo đơn hàng thủ công
// @Description Người dùng tạo đơn hàng với thông tin tùy chỉnh (không qua giỏ hàng)
// @Tags Orders
// @Accept json
// @Produce json
// @Param input body services.CreateOrderInput true "Thông tin đơn hàng"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string
// @Router /orders/manual [post]
// @Security BearerAuth
func CreateOrder(c *gin.Context) {
    var input services.CreateOrderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("user_id")
    order, err := services.CreateOrder(userID, input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.First(&user, userID).Error; err == nil {
        subject := fmt.Sprintf("Xác nhận đơn hàng #%d", order.ID)
        body := fmt.Sprintf(`
            <h3>Chào %s,</h3>
            <p>Đơn hàng của bạn đã được tạo thành công.</p>
            <ul>
              <li><strong>Tổng tiền:</strong> %.2f VND</li>
              <li><strong>Trạng thái:</strong> %s</li>
            </ul>
        `, user.Name, order.Total, order.Status)
        go config.SendMail(user.Email, subject, body)
    }

    c.JSON(http.StatusOK, order)
}

// @Summary Tạo đơn hàng từ giỏ hàng
// @Description Tạo đơn hàng dựa trên các sản phẩm trong giỏ hàng
// @Tags Orders
// @Produce json
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string
// @Router /orders/from-cart [post]
// @Security BearerAuth
func CreateOrderFromCart(c *gin.Context) {
    userID := c.GetUint("user_id")
    fmt.Println(" userID:", userID)

    order, err := services.CreateOrderFromCart(userID)
    if err != nil {
        fmt.Println(" Lỗi từ CreateOrderFromCart service:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, order)
}

// @Summary Lấy đơn hàng của người dùng hiện tại
// @Description Trả về danh sách đơn hàng của người dùng đã đăng nhập
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]string
// @Router /orders/user [get]
// @Security BearerAuth
func GetUserOrders(c *gin.Context) {
    userID := c.GetUint("user_id")
    orders, err := services.GetOrdersByUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}

// @Summary Lấy tất cả đơn hàng
// @Description Trả về toàn bộ đơn hàng trong hệ thống (chỉ dành cho admin)
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
// @Security BearerAuth
func GetAllOrders(c *gin.Context) {
    orders, err := services.GetAllOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}