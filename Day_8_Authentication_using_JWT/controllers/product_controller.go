package controllers

import (
	"jwt_authentication/config"
	"jwt_authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}

	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product added",
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"message":  "all products",
		"products": products,
	})
}
