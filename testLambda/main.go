package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	MonthNumber uint
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	if event.MonthNumber > 12 {
		return "", fmt.Errorf("the month number is greater than 12 and we got :%d", event.MonthNumber)
	}
	month := time.Month(int(event.MonthNumber))
	return month.String(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
