package main

import (
	"os"

	"github.com/dhanavadh/todol/config"
	"github.com/dhanavadh/todol/models"
	"github.com/dhanavadh/todol/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()
	routes.SetUpRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
