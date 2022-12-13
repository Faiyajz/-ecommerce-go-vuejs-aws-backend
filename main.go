package main

import (
	"fmt"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })
	// r.Run(":8000") // listen and serve on 0.0.0.0:8080
	fmt.Println(shippingPrice(9))
}

func shippingPrice(numberKg uint) int {
	if numberKg <= 10 {
		return 10
	} else if numberKg <= 20 {
		return 25
	} else {
		return 50
	}
}
