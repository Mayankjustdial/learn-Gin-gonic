package main

import "github.com/gin-gonic/gin"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 🔹 Step 2: Add Validation Rules
// binding:"required" → field is mandatory
// binding:"email" → must be valid email
// gte=18 → age ≥ 18, lte=100 → age ≤ 100

type RegisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=18,lte=100"`
}

type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type FeedbackRequest struct {
	Message string `json:"message" binding:"required"`
	Rating  int    `json:"rating" binding:"required,gte=1,lte=5"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello home routes",
		})
	})

	// 🔹 Step 1: Binding JSON to Struct

	router.POST("/user", func(c *gin.Context) {
		var user User

		// ShouldBindJSON maps incoming JSON to the User struct.
		// If the request JSON is malformed or missing fields, it returns an error.

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "USer data received",
			"user":    user,
		})

	})

	router.POST("/register", func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Registration successful",
			"data":    req,
		})
	})

	// ***********Practive Challenge****************
	// 1. 	POST /login
	// Accept JSON: { "email": "...", "password": "..." }
	// Email must be valid and required
	// Password must be min 6 chars (hint: min=6)

	router.POST("/login", func(c *gin.Context) {
		var data Credentials
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"err": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message": "User logged in Succesfully",
			"data":    data.Email,
		})
	})

	// 2.	POST /feedback
	// Accept JSON: { "message": "...", "rating": 1-5 }
	// Validate that rating is gte=1,lte=5
	// Message should be required

	router.POST("feedback", func(c *gin.Context) {
		var data FeedbackRequest

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"err": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Feedback Submitted ",
			"rating":  data,
		})
	})

	router.Run(":8080")

}
