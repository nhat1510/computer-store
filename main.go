package main

import (
    "computer-store/config"
    "computer-store/routes"
    "computer-store/seed"
    
)

func main() {
    config.ConnectDB()
    config.ConnectElasticSearch()
    seed.SeedAdminUser()
    r := routes.SetupRouter()
    r.Run(":8080")
}

