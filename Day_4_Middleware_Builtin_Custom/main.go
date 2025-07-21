package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		fmt.Printf("Mayank Incoming request: %s %s\n", method, path)
		c.Next() //continue to next handler
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "Bearer mysecrettoken" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

func TimerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start time
		start := time.Now()
		// lets the request proceed
		c.Next()
		// end time
		duration := time.Since(start)
		log.Printf("Request %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}

func HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_key := c.GetHeader("x-API_KEY")
		if api_key != "123456" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized- Invalid API key",
			})
			return
		}
		c.Next()
	}
}

func main() {
	// router := gin.Default() //Includes Logger and Recovery middlewares
	// Logger:
	// Logs each incoming request to the terminal.
	// Recovery:
	// Recovers from any panic in your handlers and prevents server crashes.
	// router

	router := gin.New() //If we are  using gin.New(), you can add middleware manually
	router.Use(gin.Logger(), gin.Recovery())

	//  CUSTOM middleware (Global)
	// create a middleware that logs the request method and path:
	router.Use(CustomLogger()) // Apply Globally - sabhi routes pr chlega

	// challenge 1
	router.GET("/user", TimerMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello middlewares",
		})
	})

	// Agar hum sirf kuch routes pe middleware lagana chahte ho
	router.GET("/secured", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Secure route assessed"})
	})

	// we can attach middleware to specific route groups:
	authorized := router.Group("/admin")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to admin dashboard"})
		})
	}

	// challenge 1
	// 	Create a request-timer middleware
	// Logs how long a request took to execute
	router.GET("/fasr", TimerMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello middlewares",
		})
	})

	// Challenge 2
	// 	Route: /profile
	// Only allow if header X-API-KEY=123456
	protected := router.Group("/password-protected")
	protected.Use(HeaderMiddleware())
	{
		router.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "welcome to profile",
			})
		})
	}
	router.Run(":8080")
}
