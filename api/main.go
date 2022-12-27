package main

import (
	"backend/internal/server"
	"backend/internal/storage"
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	allowedOrigin, found := os.LookupEnv("ALLOWED_ORIGIN")
	if !found {
		log.Fatalf("environement variable not found")
	}

	dynamoStorage, err := storage.NewDynamo("ecom-dev")
	if err != nil {
		log.Fatalf("impossible to create storage interface: #{err}")
	}

	myServer, err := server.New(server.Config{
		Port:          9000,
		AllowedOrigin: allowedOrigin,
		Storage:       dynamoStorage,
	})

	if err != nil {
		log.Fatalf("impossible to create the server: %s", err)
	}

	ginLambda = ginadapter.New(myServer.Engine)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
