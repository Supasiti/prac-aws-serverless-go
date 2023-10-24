package store

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dbclient "github.com/supasiti/prac-aws-serverless-go/internal/dynamodb"
	"github.com/supasiti/prac-aws-serverless-go/internal/pkg/json"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --with-expecter=true --name Store

type Store interface {
	user.Store
}

type store struct {
	client      dbclient.DbClient
	tableName   string
	idGenerator user.IDGenerator
}

func NewStore(client dbclient.DbClient, tableName string) (Store, error) {
	store := &store{
		client:      client,
		tableName:   tableName,
		idGenerator: idGenerator,
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
		log.Printf("store.GetUser client.GetItem error: %+v", err)
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

// CreateUser accepts User struct, save to dynamodb and return a user with new id
func (s *store) CreateUser(ctx context.Context, params *user.CreateUserParams) (*user.User, error) {
	result := user.CreateUser(params, s.idGenerator)
	userItem := result.ToUserItem()

	av, err := attributevalue.MarshalMap(userItem)
	if err != nil {
		log.Printf("store.CreateUser error: %+v", err)
		return nil, err
	}

	cmd := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(s.tableName),
	}

	_, err = s.client.PutItem(ctx, cmd)
	if err != nil {
		log.Printf("store.CreateUser client.PutItem error: %+v", err)
		return nil, err
	}

	return result, nil
}

func idGenerator() int {
	return int(time.Now().UnixMilli())
}
