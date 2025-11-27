package main

import (
	"os"

	"todol/config"
	"todol/models"
	"todol/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Todo{}, &models.User{})

	r := gin.Default()
	routes.SetUpRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
