package server

import (
	"fmt"
	"log"
	"net/http"

	"backend/internal/product"
	"backend/internal/storage"

	"github.com/Rhymond/go-money"
	uuid "github.com/satori/go.uuid"

	"backend/internal/category"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine        *gin.Engine
	port          uint
	allowedOrigin string
	storage       storage.Storage
}

type Config struct {
	Port          uint
	AllowedOrigin string
	Storage       storage.Storage
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		Engine:        engine,
		port:          config.Port,
		allowedOrigin: config.AllowedOrigin,
		storage:       config.Storage,
	}
	engine.Use(server.CORSMiddleware, server.MiddlewareServerModel, server.CheckRequest)
	engine.GET("/categories", server.Categories)
	engine.GET("/products", server.Products)
	engine.POST("/admin/products", server.CreateProduct)
	return server, nil
}

func (server *Server) Run() error {
	return server.Engine.Run(fmt.Sprintf(":%d", server.port))
}

func (server Server) MiddlewareServerModel(ctx *gin.Context) {
	ctx.Header("X-Server-Model", "Gin")
}

func (server Server) CORSMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", server.allowedOrigin)
}

func (server Server) CheckRequest(ctx *gin.Context) {
	authValue := ctx.GetHeader("Authorization")

	if authValue != "ABC" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
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
	ctx.JSON(200, products)
}

func (server Server) CreateProduct(ctx *gin.Context) {
	var productToAdd product.Product
	err := ctx.BindJSON(&productToAdd)

	if err != nil {
		log.Printf("error while binding JSON: %s \n", err)
		return
	}

	productToAdd.ID = uuid.NewV4().String()

	err = server.storage.CreateProduct(productToAdd)
	if err != nil {
		log.Printf("error occured while saving the product: %s \n", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "impossible to persist produtct"})
		return
	}
	ctx.JSON(http.StatusOK, productToAdd)
}
