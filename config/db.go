package config

import (
	"fmt"
	"log"
	"os"
	"computer-store/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(" Không thể load file .env")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" Kết nối database thất bại:", err)
	}

	fmt.Println(" Kết nối database thành công!")

	if err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Review{},
		&models.CartItem{},
		&models.News{},
	); err != nil {
		log.Fatal(" AutoMigrate thất bại:", err)
	}

	DB = db
}
