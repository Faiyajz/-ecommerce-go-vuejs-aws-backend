package main

import (
	"backend/internal/product"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, product.Product{
			ID:               "868",
			Name:             "Handbook",
			Description:      "July, 2022",
			PriceVATExcluded: money.New(1000,"USD"),
			VAT:              money.New(1200, "USD"),
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
