package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/supasiti/prac-aws-serverless-go/internal/dynamodb"
	"github.com/supasiti/prac-aws-serverless-go/internal/handler/get_user"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
)

type handler = func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

func main() {
	tableName := os.Getenv("USER_TABLE_NAME")
	if len(tableName) == 0 {
		log.Print("Missing table name")
		return
	}

	dbclient := dynamodb.GetDbClient()
	store, err := store.NewStore(dbclient, tableName)
	if err != nil {
		log.Printf("Unable to create connection to dynamodb: %+v", err)
		return
	}

	handler := createHandler(store)
	lambda.Start(handler)
}

func createHandler(store store.Store) handler {
	handler := getuser.NewHandler(store)

	return func(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		res := handler.GetUser(ctx, req)
		return res, nil
	}
}
