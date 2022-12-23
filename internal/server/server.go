package server

import (
	"fmt"
	"log"
	"net/http"

	"backend/internal/product"

	"github.com/Rhymond/go-money"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"backend/internal/category"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine        *gin.Engine
	port          uint
	allowedOrigin string
}

type Config struct {
	Port          uint
	AllowedOrigin string
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		Engine:        engine,
		port:          config.Port,
		allowedOrigin: config.AllowedOrigin,
	}
	engine.Use(server.CORSMiddleware, server.MiddlewareServerModel, server.CheckRequest)
	engine.GET("/categories", server.Categories)
	engine.GET("/products", server.Products)
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

	awsSession, err := session.NewSession()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	dynamodbClient := dynamodb.New(awsSession)

	tableName := "ecommerce-dev"
	item := make(map[string]*dynamodb.AttributeValue)

	item["PK"] = &dynamodb.AttributeValue{
		S: aws.String("test"),
	}
	item["SK"] = &dynamodb.AttributeValue{
		S: aws.String("test2"),
	}
	item["foo"] = &dynamodb.AttributeValue{
		S: aws.String("bar"),
	}

	output, err := dynamodbClient.PutItem(&dynamodb.PutItemInput{
		Item:      nil,
		TableName: &tableName,
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error - database query error"})
		return
	}

	log.Println(output)

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
