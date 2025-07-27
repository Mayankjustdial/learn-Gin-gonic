package controllers

import (
	"jwt_authentication/config"
	"jwt_authentication/models"
	"jwt_authentication/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed

	config.DB.Create(&user)
	c.JSON(http.StatusOK, user)

}

func Login(c *gin.Context) {
	var input_auth models.User
	var user_found models.User

	if err := c.ShouldBindJSON(&input_auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := config.DB.Where("username = ?", input_auth.Username).Find(&user_found).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials"})
		return
	}
	if !utils.VerifyPassword(input_auth.Password, user_found.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_found.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})

	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully loggedin",
		"token":   tokenString,
	})

}
