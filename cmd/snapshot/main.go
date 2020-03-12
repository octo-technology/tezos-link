package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    "log"
)

type MyEvent struct {
    Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    log.Print("Snapshot export event received.")

    return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
    lambda.Start(HandleRequest)
}
