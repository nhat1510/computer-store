package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConnectRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Không thể load file .env:", err)
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
	log.Println("❌ Không thể kết nối Redis:", err)
} else {
	log.Println("✅ Đã kết nối Redis thành công!")
}
}