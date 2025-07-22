package main

import (
	"environment-veriables/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	port := config.GetEnv("PORT", "8080")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Setup environment veriables",
		})
	})

	router.Run(":" + port)
}
