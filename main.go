package main

import (
	"backend/internal/category"
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
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		ctx.JSON(200, products)
	})

	r.GET("/categories", func(ctx *gin.Context) {
		categories := []category.Category{
			{
				ID:          "1",
				Name:        "HardCover",
				Description: "A Hard Copy of the Handbook",
			},
			{
				ID:          "2",
				Name:        "Digital Version",
				Description: "A Soft Copy of the Handbook",
			},
		}
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		ctx.JSON(200, categories)
	})
	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}
