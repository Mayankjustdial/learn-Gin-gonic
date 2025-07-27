package routes

import (
	"jwt_authentication/controllers"
	"jwt_authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/signup", controllers.Signup)
		api.POST("/login", controllers.Login)

		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.POST("/product", controllers.AddProduct)
			protected.GET("/products", controllers.GetProducts)
		}

	}

}
