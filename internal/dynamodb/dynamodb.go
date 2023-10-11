package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dbClient *dynamodb.Client
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --with-expecter=true --name DbClient

type DbClient interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func NewDbClient(optFns ...func(*dynamodb.Options)) (*dynamodb.Client, error) {
	// config := aws.Config{
	// 	RetryMaxAttempts: 3,
	// 	RetryMode:        aws.RetryModeStandard,
	// }
	//
	// for _, option := range options {
	// 	option(&config)
	// }

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRetryMaxAttempts(3),
		config.WithRetryMode(aws.RetryModeStandard),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg, optFns...), nil
}

func GetDbClient(optFns ...func(*dynamodb.Options)) (*dynamodb.Client, error) {
	if dbClient == nil {
		client, err := NewDbClient(optFns...)
		if err != nil {
			return nil, err
		}
		dbClient = client
	}

	return dbClient, nil
}
