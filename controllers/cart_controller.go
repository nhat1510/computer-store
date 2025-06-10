package controllers

import (
    "computer-store/services"
    "log"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
    var input services.AddToCartInput
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println("❌ Lỗi bind JSON trong AddToCart:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("user_id")
    log.Printf("🛒 AddToCart: user_id=%d, product_id=%d, quantity=%d\n", userID, input.ProductID, input.Quantity)

    cart, err := services.AddToCart(userID, input)
    if err != nil {
        log.Println("❌ Lỗi thêm vào giỏ hàng:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Println("✅ Đã thêm vào giỏ hàng:", cart)
    c.JSON(http.StatusOK, cart)
}

func GetCart(c *gin.Context) {
    userID := c.GetUint("user_id")
    log.Printf("📦 Lấy giỏ hàng cho userID: %d\n", userID)

    cartItems, err := services.GetCartItems(userID)
    if err != nil {
        log.Println("❌ Lỗi khi lấy giỏ hàng:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    log.Printf("✅ Giỏ hàng có %d sản phẩm\n", len(cartItems))
    c.JSON(http.StatusOK, cartItems)
}

func UpdateCartItem(c *gin.Context) {
    userID := c.GetUint("user_id")
    productIDStr := c.Param("product_id")
    productID, err := strconv.ParseUint(productIDStr, 10, 64)
    if err != nil {
        log.Println("❌ Lỗi chuyển đổi product_id:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    var input services.UpdateCartInput
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println("❌ Lỗi bind JSON trong UpdateCartItem:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Printf("🔁 Cập nhật giỏ hàng: user_id=%d, product_id=%d, new_quantity=%d\n", userID, productID, input.Quantity)

    err = services.UpdateCartItem(userID, uint(productID), input.Quantity)
    if err != nil {
        log.Println("❌ Lỗi cập nhật giỏ hàng:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Println("✅ Cập nhật số lượng thành công")
    c.JSON(http.StatusOK, gin.H{"message": "Cập nhật thành công"})
}

func RemoveFromCart(c *gin.Context) {
    userID := c.GetUint("user_id")
    productIDStr := c.Param("product_id")
    productID, err := strconv.ParseUint(productIDStr, 10, 64)
    if err != nil {
        log.Println("❌ Lỗi chuyển đổi product_id:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    log.Printf("🗑️ Xoá sản phẩm khỏi giỏ hàng: user_id=%d, product_id=%d\n", userID, productID)

    err = services.RemoveCartItem(userID, uint(productID))
    if err != nil {
        log.Println("❌ Lỗi xoá sản phẩm:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    log.Println("✅ Đã xoá khỏi giỏ hàng")
    c.JSON(http.StatusOK, gin.H{"message": "Đã xoá sản phẩm khỏi giỏ hàng"})
}
