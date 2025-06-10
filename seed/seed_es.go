package seed

import (
	"computer-store/config"
	"computer-store/models"
	"bytes"
	"encoding/json"
	"fmt"
)

func IndexAllProductsToElastic() error {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return err
	}

	for _, product := range products {
		body, _ := json.Marshal(product)
		res, err := config.ES.Index("products", bytes.NewReader(body))
		if err != nil {
			return err
		}
		defer res.Body.Close()
	}

	fmt.Println(" Đã index toàn bộ sản phẩm vào ElasticSearch")
	return nil
}
