package controllers

import (
	"computer-store/config"
	"computer-store/models"
	"computer-store/services"
	"computer-store/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

// CreateProduct tạo mới một sản phẩm mới từ form-data
// @Summary Tạo sản phẩm
// @Tags Products
// @Accept mpfd
// @Produce json
// @Param name formData string true "Tên sản phẩm"
// @Param description formData string true "Mô tả"
// @Param price formData number true "Giá"
// @Param stock formData int true "Số lượng"
// @Param category_id formData int true "ID danh mục"
// @Param image formData file true "Ảnh sản phẩm"
// @Success 200 {object} models.Product
// @Failure 400,500 {object} map[string]string
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	stockStr := c.PostForm("stock")
	categoryIDStr := c.PostForm("category_id")

	price, _ := strconv.ParseFloat(priceStr, 64)
	stock, _ := strconv.Atoi(stockStr)
	categoryID, _ := strconv.Atoi(categoryIDStr)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu file ảnh sản phẩm"})
		return
	}

	filename := utils.GenerateFileName(file.Filename)
	savePath := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu ảnh"})
		return
	}

	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Image:       filename,
		CategoryID:  uint(categoryID),
	}

	created, err := services.CreateProductFromForm(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, created)
}

// UpdateProduct cập nhật thông tin sản phẩm
// @Summary Cập nhật sản phẩm
// @Tags Products
// @Accept mpfd
// @Produce json
// @Param id path string true "ID sản phẩm"
// @Param name formData string false "Tên sản phẩm"
// @Param description formData string false "Mô tả"
// @Param price formData number false "Giá"
// @Param stock formData int false "Số lượng"
// @Param category_id formData int false "ID danh mục"
// @Param image formData file false "Ảnh mới (nếu có)"
// @Success 200 {object} models.Product
// @Failure 400,404,500 {object} map[string]string
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	name := c.PostForm("name")
	description := c.PostForm("description")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	categoryID, _ := strconv.Atoi(c.PostForm("category_id"))

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
		return
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock
	product.CategoryID = uint(categoryID)

	file, err := c.FormFile("image")
	if err == nil {
		filename := utils.GenerateFileName(file.Filename)
		savePath := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, savePath); err == nil {
			product.Image = filename
		}
	}

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi cập nhật sản phẩm"})
		return
	}

	config.RedisClient.Del(config.Ctx, "products:all") // 🧹 Xoá cache sau cập nhật

	c.JSON(http.StatusOK, product)
}

// GetAllProducts trả về toàn bộ danh sách sản phẩm
// @Summary Danh sách sản phẩm
// @Tags Products
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// DeleteProduct xoá sản phẩm theo ID
// @Summary Xoá sản phẩm
// @Tags Products
// @Param id path string true "ID sản phẩm"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đã xóa sản phẩm"})
}

// GetProductByID trả về chi tiết sản phẩm theo ID
// @Summary Xem chi tiết sản phẩm
// @Tags Products
// @Produce json
// @Param id path string true "ID sản phẩm"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
