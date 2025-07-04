// @title Computer Store API
// @version 1.0
// @description API cho hệ thống bán hàng máy tính
// @termsOfService http://localhost:8080/
// @contact.name Developer
// @contact.email you@example.com
// @license.name MIT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
    "computer-store/config"
    "computer-store/routes"
    "computer-store/seed"
    
)

func main() {
    config.ConnectDB()
    config.ConnectElasticSearch()
    config.ConnectRedis()
    seed.SeedAdminUser()
    
    r := routes.SetupRouter()
    r.Run(":8080")
}

