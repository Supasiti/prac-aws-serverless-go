package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	dbclient "github.com/supasiti/prac-aws-serverless-go/internal/dynamodb"
	"github.com/supasiti/prac-aws-serverless-go/internal/pkg/json"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
)

type Store interface {
	user.Store
}

type store struct {
	client    dbclient.DbClient
	tableName string
}

func NewStore(client dbclient.DbClient, tableName string) (Store, error) {
	store := &store{
		client:    client,
		tableName: tableName,
	}

	return store, nil
}

func (s *store) GetUser(ctx context.Context, userID int) (*user.User, error) {
	key, err := user.GetKey(userID)
	if err != nil {
		return nil, err
	}

	cmd := &dynamodb.GetItemInput{
		TableName: aws.String(s.tableName),
		Key:       key,
	}
	log.Printf("store.GetUser command: %s", json.ToInlineJSON(cmd))

	data, err := s.client.GetItem(ctx, cmd)
	if err != nil {
		log.Printf("store.GetUser client.GetItem error: %s", json.ToInlineJSON(err))
		return nil, err
	}

	log.Printf("store.GetUser client.GetItem data: %s", json.ToInlineJSON(data))
	if data.Item == nil {
		return nil, user.ErrUserNotFound
	}

	res := user.User{}

	err = attributevalue.UnmarshalMap(data.Item, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
