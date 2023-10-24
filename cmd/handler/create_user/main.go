package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/supasiti/prac-aws-serverless-go/internal/dynamodb"
	createuser "github.com/supasiti/prac-aws-serverless-go/internal/handler/create_user"
	"github.com/supasiti/prac-aws-serverless-go/internal/pkg/json"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
)

type handler = func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

func main() {
	tableName := os.Getenv("USER_TABLE_NAME")
	if len(tableName) == 0 {
		log.Print("Missing table name")
		return
	}

	dbclient, err := dynamodb.GetDbClient()
	if err != nil {
		log.Printf("Unable to create connection to dynamodb: %s", json.ToInlineJSON(err))
		return
	}

	store, err := store.NewStore(dbclient, tableName)
	if err != nil {
		log.Printf("Unable to create store: %s", json.ToInlineJSON(err))
		return
	}

	h := createHandler(store)
	lambda.Start(h)
}

func createHandler(store store.Store) handler {
	h := createuser.NewHandler(store)

	return func(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		res := h.CreateUser(ctx, req)
		return res, nil
	}
}
