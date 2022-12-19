package server

import (
	"fmt"
	"net/http"

	"backend/internal/product"

	"github.com/Rhymond/go-money"

	"backend/internal/category"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	port   uint
}

type Config struct {
	Port uint
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		engine: engine,
		port:   config.Port,
	}
	engine.GET("/categories", server.Categories)
	engine.GET("/products", server.Products)
	return server, nil
}

func (server *Server) Run() error {
	return server.engine.Run(fmt.Sprintf(":%d", server.port))
}

func (server *Server) Categories(ctx *gin.Context) {
	categories := []category.Category{
		{
			ID:          "42",
			Name:        "Plushies",
			Description: "kdsjdjsidjisdj",
		},
		{
			ID:          "43",
			Name:        "T-Shirts",
			Description: "kdsjdjsidjisdj",
		},
	}
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.JSON(http.StatusOK, categories)
}

func (server *Server) Products(ctx *gin.Context) {
	products := []product.Product{
		{
			ID:               "42",
			Name:             "Test",
			Description:      "This is my product",
			PriceVATExcluded: money.New(100, "EUR"),
			VAT:              money.New(200, "EUR"),
		},
		{
			ID:               "43",
			Name:             "Test",
			Description:      "This is my product",
			PriceVATExcluded: money.New(100, "EUR"),
			VAT:              money.New(250, "EUR"),
		},
		{
			ID:               "44",
			Name:             "Test",
			Description:      "This is my product",
			PriceVATExcluded: money.New(189, "EUR"),
			VAT:              money.New(200, "EUR"),
		},
	}
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.JSON(200, products)
}
