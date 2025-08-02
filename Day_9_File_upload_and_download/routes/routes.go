package routes

import (
	"file_upload/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine) {
	r.POST("/upload", controllers.UploadFile)
	r.GET("/download/:filename", controllers.DownloadFile)
}
