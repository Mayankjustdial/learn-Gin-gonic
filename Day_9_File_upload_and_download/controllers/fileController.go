package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Define path to save
	// 	This defines where on your server the uploaded file will be saved.
	// "uploads" is your destination folder
	// file.Filename is the original file name
	// filepath.Join() safely creates uploads/filename.ext path, cross-platform
	uploadPath := filepath.Join("uploads", file.Filename)

	// Save file to disk
	// Gin saves the uploaded file to the path you created above
	// If it fails (e.g., folder doesn't exist, permission denied), it returns an error

	err = c.SaveUploadedFile(file, uploadPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File save failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "File uploaded successfully",
		"file_name":  file.Filename,
		"file_size":  file.Size,
		"uploadPath": uploadPath,
	})
}

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("uploads", filename)

	// Let client download the file
	c.FileAttachment(filePath, filename)
}
