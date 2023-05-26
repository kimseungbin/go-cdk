package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

// create lambda function returning hello world and 200 status.

type MyEvent struct {
	Name string `json:"name"`
}

// handler handles incoming lambda request.
func handler(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello, %s", name.Name), nil
}
func main() {
	lambda.Start(handler)
}
