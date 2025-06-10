package controllers

import (
    "computer-store/config"
    "computer-store/models"
    "computer-store/services"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

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


func GetUserOrders(c *gin.Context) {
    userID := c.GetUint("user_id")
    orders, err := services.GetOrdersByUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}

func GetAllOrders(c *gin.Context) {
    orders, err := services.GetAllOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}