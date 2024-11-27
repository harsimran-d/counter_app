package main

import (
	"log"
	"net/http"
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

	router.GET("/api/increment", func(c *gin.Context) {
		globalCount.mu.Lock()
		defer globalCount.mu.Unlock()
		globalCount.value++
		c.JSON(200, globalCount.value)
	})

	router.StaticFS("./public", http.Dir(""))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})
	log.Printf("Server started on port: %v\n", port)
	log.Fatal(router.Run(address))
}
