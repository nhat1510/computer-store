package controllers

import (
    "net/http"
    "computer-store/models"
    "computer-store/services"
	"computer-store/config"
	"computer-store/utils"
	"path/filepath"
	"strconv"
    "github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
    products, err := services.GetAllProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    if err := services.DeleteProduct(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đã xóa sản phẩm"})
}

func GetProductByID(c *gin.Context) {
    id := c.Param("id")
    product, err := services.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, product)
}
