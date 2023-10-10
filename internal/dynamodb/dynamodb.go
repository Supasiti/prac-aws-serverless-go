package dynamodb

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dbClient *dynamodb.Client
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --with-expecter=true --name DbClient

type DbClient interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

type DbClientOption = func(*aws.Config)

func NewDbClient(options ...DbClientOption) *dynamodb.Client {
	region := os.Getenv("AWS_REGION")

	config := aws.Config{
		Region:           *aws.String(region),
		RetryMaxAttempts: 3,
		RetryMode:        aws.RetryModeStandard,
	}

	for _, option := range options {
		option(&config)
	}

	return dynamodb.NewFromConfig(config)
}

func GetDbClient(options ...DbClientOption) *dynamodb.Client {
	if dbClient == nil {
		dbClient = NewDbClient(options...)
	}

	return dbClient
}

func WithRegion(region string) DbClientOption {
	return func(c *aws.Config) {
		c.Region = *aws.String(region)
	}
}

func WithRetryMaxAttempts(num int) DbClientOption {
	return func(c *aws.Config) {
		c.RetryMaxAttempts = num
	}
}

func WithRetryMode(mode aws.RetryMode) DbClientOption {
	return func(c *aws.Config) {
		c.RetryMode = mode
	}
}

func WithConfig(dbConfig aws.Config) DbClientOption {
	return func(c *aws.Config) {
		c.Region = dbConfig.Region
		c.DefaultsMode = dbConfig.DefaultsMode
		c.RuntimeEnvironment = dbConfig.RuntimeEnvironment
		c.HTTPClient = dbConfig.HTTPClient
		c.Credentials = dbConfig.Credentials
		c.APIOptions = dbConfig.APIOptions
		c.Logger = dbConfig.Logger
		c.ClientLogMode = dbConfig.ClientLogMode
		c.AppID = dbConfig.AppID

	}
}
