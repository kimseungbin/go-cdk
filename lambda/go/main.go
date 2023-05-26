package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handler handles incoming lambda request.
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	fmt.Printf("request: %s\n", request.Body)

	response = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       fmt.Sprintf("Good Morning, CDK! You've hit %s\n", request.Path),
	}

	return
}

func main() {
	lambda.Start(handler)
}
