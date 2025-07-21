package routes

import (
	"learn-modular_Project_structure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
	}
}
