package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	address := ":" + port

	globalCount := GloblaCount{}

	router := gin.Default()
	router.Use(cors())
	router.GET("/increment", func(c *gin.Context) {
		globalCount.mu.Lock()
		defer globalCount.mu.Unlock()
		globalCount.value++
		c.JSON(200, globalCount.value)
	})
	log.Printf("Server started on port: %v\n", port)
	log.Fatal(router.Run(address))
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
