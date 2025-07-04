package controllers

import (
    "computer-store/services"
    "log"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// @Summary Thêm sản phẩm vào giỏ hàng
// @Description Người dùng thêm một sản phẩm vào giỏ hàng của họ
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body services.AddToCartInput true "Dữ liệu sản phẩm cần thêm"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /cart [post]
// @Security BearerAuth
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

// @Summary Lấy giỏ hàng của người dùng
// @Description Lấy danh sách sản phẩm trong giỏ hàng của người dùng hiện tại
// @Tags Cart
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /cart [get]
// @Security BearerAuth
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

// @Summary Cập nhật số lượng sản phẩm trong giỏ hàng
// @Description Cập nhật số lượng của một sản phẩm trong giỏ hàng theo product_id
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id path int true "ID sản phẩm"
// @Param input body services.UpdateCartInput true "Số lượng mới"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /cart/{product_id} [put]
// @Security BearerAuth
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

// @Summary Xóa sản phẩm khỏi giỏ hàng
// @Description Xóa một sản phẩm ra khỏi giỏ hàng của người dùng
// @Tags Cart
// @Produce json
// @Param product_id path int true "ID sản phẩm"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/{product_id} [delete]
// @Security BearerAuth
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
