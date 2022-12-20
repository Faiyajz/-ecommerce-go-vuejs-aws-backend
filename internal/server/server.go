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
	Engine *gin.Engine
	port   uint
}

type Config struct {
	Port uint
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		Engine: engine,
		port:   config.Port,
	}
	engine.GET("/categories", server.Categories)
	engine.GET("/products", server.Products)
	return server, nil
}

func (server *Server) Run() error {
	return server.Engine.Run(fmt.Sprintf(":%d", server.port))
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

	twoEuro := money.New(200, "EUR")
	fourEuro := money.New(400, "EUR")

	products := []product.Product{
		{
			ID:               "42",
			Name:             "Handbook Digital Version",
			ShortDescription: "July, 2022",
			Description:      "This is my product",
			PriceVATExcluded: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			VAT: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			TotalPrice: product.Amount{
				Money:   fourEuro,
				Display: fourEuro.Display(),
			},
			Image: "https://cdn.pixabay.com/photo/2016/03/31/20/51/book-1296045_1280.png",
		},
		{
			ID:               "43",
			Name:             "Test",
			ShortDescription: "New",
			Description:      "This is my product",
			PriceVATExcluded: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			VAT: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			TotalPrice: product.Amount{
				Money:   fourEuro,
				Display: fourEuro.Display(),
			},
			Image: "https://cdn.pixabay.com/photo/2014/09/08/05/06/book-438935_1280.png",
		},
		{
			ID:               "44",
			Name:             "Test",
			Description:      "This is my product",
			ShortDescription: "New",
			PriceVATExcluded: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			VAT: product.Amount{
				Money:   twoEuro,
				Display: twoEuro.Display(),
			},
			TotalPrice: product.Amount{
				Money:   fourEuro,
				Display: fourEuro.Display(),
			},
			Image: "https://cdn.pixabay.com/photo/2012/04/12/13/54/book-30127_1280.png",
		},
	}
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.JSON(200, products)
}
