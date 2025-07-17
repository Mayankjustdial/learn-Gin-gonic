package main

import "github.com/gin-gonic/gin"

/*

Understand and implement:

Path parameters (/user/:id)

Query parameters (?q=search)

Route grouping (/api/v1/...)

*/

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})

	})

	// ***********Path Parameters (Dynamic Segments)***************
	// 	:id is a dynamic value in the URL.
	// c.Param("id") fetches the value from the URL.
	// Example-------
	// 	Request → GET /user/123
	// Response → {"user_id": "123"}

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"user_id": id})
	})

	// *********** Query Parameters *************************
	// 	Query parameters come after ? in the URL.
	// c.Query("q") gets the value from the query string.
	// Example--------
	// Request → GET /search?q=golang
	// Response → {"search_term": "golang"}

	router.GET("/search", func(c *gin.Context) {
		query := c.Query("q") // or c.DefaultQuery("q", "default")
		c.JSON(200, gin.H{"search_term": query})
	})

	// ************Route Grouping ***************************
	// 	Groups keep related routes organized.
	// Useful for versioning APIs or organizing by module.
	v1 := router.Group("api/v1")
	{
		v1.GET("/books", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "All books"})
		})

		v1.GET("/books/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"book_id": id})
		})
	}

	//*********************** Day two challenges***************************
	// 	1. Create a route: /product/:category/:id
	// Return category and id in the response.

	router.GET("/product/:category/:id", func(c *gin.Context) {
		category := c.Param("category")
		product_id := c.Param("id")

		c.JSON(200, gin.H{
			"category": category,
			"id":       product_id,
		})

	})

	// 	Query parameter filter:
	// Route: /filter
	// Use ?min=10&max=100 and return both.

	router.GET("/product/filter", func(c *gin.Context) {
		min := c.Query("min")
		max := c.Query("max")
		c.JSON(200, gin.H{
			"min": min,
			"max": max,
		})
	})

	// 	3.Group route /api/v1/profile/:username
	// Should return: { "username": "..." }
	v2 := router.Group("/api/v2")
	{
		v2.GET("/profile/:username", func(c *gin.Context) {
			username := c.Param("username")
			c.JSON(200, gin.H{
				"username": username,
			})
		})
	}

	router.Run(":8080")

}
