package main

import (
	"jwt_authentication/config"
	"jwt_authentication/models"
	"jwt_authentication/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	port := config.GetEnv("PORT", "8080")
	router := gin.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Product{})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "JWT Authentication",
		})
	})
	routes.RegisterRoutes(router)

	router.Run(":" + port)
}
