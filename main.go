package main

import (
	"backend/internal/server"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	myServer, err := server.New(server.Config{})

	if err != nil {
		log.Fatalf("impossible to create the server: %s", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/products", myServer.Products)
	r.GET("/categories", myServer.Categories)
	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}
