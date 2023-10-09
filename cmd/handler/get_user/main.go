package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("HELLO: %s", os.Getenv("HELLO"))

	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello world",
	}

	return &res, nil
}
