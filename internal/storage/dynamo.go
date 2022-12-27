package storage

import (
	"backend/internal/product"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Dynamo struct {
	tableName  string
	awsSession *session.Session
	client     *dynamodb.DynamoDB
}

func NewDynamo(tableName string) (*Dynamo, error) {
	awsSession, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("impossible to create aws session: %w", err)
	}

	dynamodbClient := dynamodb.New(awsSession)

	return &Dynamo{
		tableName:  tableName,
		awsSession: awsSession,
		client:     dynamodbClient,
	}, nil
}

func (dynamo *Dynamo) CreateProduct(product product.Product) error {

	// item := make(map[string]*dynamodb.AttributeValue)

	item, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return fmt.Errorf("impossible to marshal product: %w", err)
	}

	item["PK"] = &dynamodb.AttributeValue{
		S: aws.String("product"),
	}
	item["SK"] = &dynamodb.AttributeValue{
		S: aws.String("product.ID"),
	}

	_, err = dynamo.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &dynamo.tableName,
	})

	if err != nil {
		return fmt.Errorf("impossible to Put Item in db: %w", err)
	}

	return nil

}
