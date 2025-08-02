package main

import (
	"file_upload/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Home routes",
		})
	})

	routes.FileRoutes(router)
	router.Run(":8080")

}
