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
		products := []product.Product{
			{
				ID:               "868",
				Name:             "Handbook",
				Description:      "July, 2022",
				PriceVATExcluded: money.New(1000, "USD"),
				VAT:              money.New(1200, "USD"),
			},
			{
				ID:               "869",
				Name:             "Handbook Hardcover",
				Description:      "December, 2022",
				PriceVATExcluded: money.New(2000, "USD"),
				VAT:              money.New(2500, "USD"),
			},
			{
				ID:               "870",
				Name:             "Handbook",
				Description:      "New Edition",
				PriceVATExcluded: money.New(1500, "USD"),
				VAT:              money.New(1800, "USD"),
			},
		}
		ctx.JSON(200, products)
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
