package main

import (
	"fmt"
	"gormPostgres/config"
	"gormPostgres/models"
	"gormPostgres/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	port := config.GetEnv("PORT", "8080")
	router := gin.New()

	// Connect to Database
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Product{})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "home route",
		})
	})

	routes.RegisterRoutes(router)
	fmt.Println("✅ Server is running on PORT ::" + port)
	router.Run(":" + port)

}
