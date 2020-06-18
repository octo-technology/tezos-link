package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest execute the logic
func HandleRequest(ctx context.Context) (string, error) {
	log.Print("metrics clean starting")

	return "metric clean started.", nil
}
