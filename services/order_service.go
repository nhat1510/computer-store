package services

import (
    "computer-store/config"
    "computer-store/models"
    "errors"
    "fmt"
)

type CreateOrderInput struct {
    Items []struct {
        ProductID uint `json:"product_id"`
        Quantity  int  `json:"quantity"`
    } `json:"items"`
}

func CreateOrder(userID uint, input CreateOrderInput) (models.Order, error) {
    var total float64
    var items []models.OrderItem

    for _, i := range input.Items {
        var product models.Product
        if err := config.DB.First(&product, i.ProductID).Error; err != nil {
            return models.Order{}, errors.New("Sản phẩm không tồn tại")
        }

        if i.Quantity > product.Stock {
            return models.Order{}, fmt.Errorf("Số lượng vượt quá tồn kho: %s", product.Name)
        }

        total += float64(product.Price) * float64(i.Quantity)

        items = append(items, models.OrderItem{
            ProductID: i.ProductID,
            Quantity:  i.Quantity,
            Price:     float64(product.Price),
        })

        product.Stock -= i.Quantity
        config.DB.Save(&product)
    }

    order := models.Order{
        UserID: userID,
        Total:  total,
        Status: "pending",
        Items:  items,
    }

    if err := config.DB.Create(&order).Error; err != nil {
        return models.Order{}, errors.New("Không thể tạo đơn hàng")
    }

    return order, nil
}

func CreateOrderFromCart(userID uint) (models.Order, error) {
    var cartItems []models.CartItem
    if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
        return models.Order{}, errors.New("Không thể lấy giỏ hàng")
    }
    if len(cartItems) == 0 {
        return models.Order{}, errors.New("Giỏ hàng trống")
    }

    var orderItems []models.OrderItem
    var total float64

    for _, item := range cartItems {
        product := item.Product

        if item.Quantity > product.Stock {
            return models.Order{}, fmt.Errorf("Sản phẩm '%s' không đủ hàng", product.Name)
        }

        total += float64(product.Price) * float64(item.Quantity)

        orderItems = append(orderItems, models.OrderItem{
            ProductID: product.ID,
            Quantity:  item.Quantity,
            Price:     float64(product.Price),
        })

        product.Stock -= item.Quantity
        config.DB.Save(&product)
    }

    order := models.Order{
        UserID: userID,
        Total:  total,
        Status: "pending",
        Items:  orderItems,
    }

    if err := config.DB.Create(&order).Error; err != nil {
        return models.Order{}, errors.New("Không thể tạo đơn hàng")
    }

    config.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})
    return order, nil
}

func GetOrdersByUser(userID uint) ([]models.Order, error) {
    var orders []models.Order
    err := config.DB.
        Preload("Items.Product"). // 👈 thêm dòng này
        Where("user_id = ?", userID).
        Find(&orders).Error
    if err != nil {
        return nil, errors.New("Không thể lấy đơn hàng")
    }

    // Gắn thêm ảnh vào từng item
    for i := range orders {
        for j := range orders[i].Items {
            orders[i].Items[j].ProductImage = orders[i].Items[j].Product.Image // 👈 đảm bảo có trường này
        }
    }

    return orders, nil
}


func GetAllOrders() ([]models.Order, error) {
    var orders []models.Order
    if err := config.DB.Preload("Items").Preload("Items.Product").Find(&orders).Error; err != nil {
        return nil, errors.New("Không thể lấy danh sách đơn hàng")
    }
    return orders, nil
}

