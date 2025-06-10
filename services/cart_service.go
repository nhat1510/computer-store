package services

import (
	"computer-store/config"
	"computer-store/models"
	"errors"
	"log"
)

type AddToCartInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type UpdateCartInput struct {
	Quantity int `json:"quantity"`
}

func AddToCart(userID uint, input AddToCartInput) (models.CartItem, error) {
	log.Println("🛒 AddToCart called with:", "userID =", userID, "productID =", input.ProductID, "quantity =", input.Quantity)

	if input.ProductID == 0 || input.Quantity <= 0 {
		log.Println("❌ Invalid input - ProductID or Quantity not valid")
		return models.CartItem{}, errors.New("ProductID và Quantity không hợp lệ")
	}

	var existing models.CartItem
	err := config.DB.Preload("Product").Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&existing).Error
	if err == nil {
		log.Println("ℹ️ Product đã có trong giỏ, tăng số lượng")
		existing.Quantity += input.Quantity
		if err := config.DB.Save(&existing).Error; err != nil {
			log.Println("❌ ERROR khi cập nhật giỏ hàng:", err)
			return models.CartItem{}, err
		}
		// preload lại Product để đảm bảo dữ liệu đầy đủ
		config.DB.Preload("Product").First(&existing, existing.ID)
		return existing, nil
	}

	log.Println("🆕 Thêm sản phẩm mới vào giỏ hàng")
	cart := models.CartItem{
		UserID:    userID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}
	if err := config.DB.Create(&cart).Error; err != nil {
		log.Println("❌ ERROR khi tạo mới cart_item:", err)
		return models.CartItem{}, err
	}

	// ✅ Preload lại Product để trả về đầy đủ dữ liệu
	if err := config.DB.Preload("Product").First(&cart, cart.ID).Error; err != nil {
		log.Println("❌ ERROR khi preload product sau khi thêm:", err)
		return models.CartItem{}, err
	}

	log.Println("✅ Đã thêm vào giỏ:", cart)
	return cart, nil
}

func GetCartItems(userID uint) ([]models.CartItem, error) {
	log.Println("📦 GetCartItems for userID:", userID)
	var items []models.CartItem
	if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&items).Error; err != nil {
		log.Println("❌ ERROR khi lấy giỏ hàng:", err)
		return nil, err
	}
	log.Printf("✅ Tìm thấy %d mục trong giỏ hàng\n", len(items))
	return items, nil
}

func UpdateCartItem(userID, productID uint, quantity int) error {
	log.Println("✏️ UpdateCartItem:", "userID =", userID, "productID =", productID, "quantity =", quantity)

	var item models.CartItem
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error; err != nil {
		log.Println("❌ Không tìm thấy mục giỏ hàng:", err)
		return errors.New("Mục giỏ hàng không tồn tại")
	}

	item.Quantity = quantity
	if err := config.DB.Save(&item).Error; err != nil {
		log.Println("❌ Lỗi khi cập nhật số lượng:", err)
		return err
	}
	log.Println("✅ Cập nhật thành công")
	return nil
}

func RemoveCartItem(userID, productID uint) error {
	log.Println("🗑️ RemoveCartItem:", "userID =", userID, "productID =", productID)
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.CartItem{}).Error; err != nil {
		log.Println("❌ Lỗi khi xoá khỏi giỏ hàng:", err)
		return err
	}
	log.Println("✅ Xoá thành công khỏi giỏ hàng")
	return nil
}
