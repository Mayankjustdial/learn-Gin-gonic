package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() // Create a gin router with Logger and Recovery
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello GIn!",
		})
	})

	router.POST("/submit", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Data recieved",
		})
	})
	router.Run(":8080")

}

/*package main
This tells Go this is an executable program, not a library.

🔸 import "github.com/gin-gonic/gin"
Brings in the Gin web framework so we can use it to create routes and handle requests.

🔸 r := gin.Default()
This creates a default Gin engine.

It includes two helpful middleware by default:

Logger: logs each HTTP request to the console.

Recovery: recovers from any panics and returns a 500 instead of crashing the server.

🔸 r.GET("/ping", func(c *gin.Context) { ... })
This sets up a GET endpoint at /ping.

When a user hits http://localhost:8080/ping, this function is called.

c *gin.Context gives access to the request and allows us to respond.

🔸 c.JSON(200, gin.H{"message": "pong"})
Sends a JSON response with HTTP status 200.

gin.H is just a shorthand for map[string]interface{}.

So the response is: {"message": "pong"}.

🔸 r.Run(":8080")
Starts the HTTP server on port 8080.

You can change the port (e.g., r.Run(":3000")).

*/
