package routes

import (
	"gormPostgres/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/products", controllers.CreateProduct)
		api.GET("/products", controllers.GetProducts)
	}
}
