package user

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/supasiti/prac-aws-serverless-go/internal/pkg/json"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/table"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Store interface {
	// GetUser accepts userID and finds and returns a user associated with it
	// return error if not found
	GetUser(context.Context, int) (*User, error)
}

type User struct {
	UserID    int     `json:"userId" dynamodbav:"userId"`
	FirstName string  `json:"firstName" dynamodbav:"firstName"`
	Balance   float64 `json:"balance,omitempty" dynamodbav:"balance"`
	Email     string  `json:"email" dynamodbav:"email"`
}

func (u User) String() string {
	return json.ToJSONString(u)
}

type UserItem struct {
	User
	table.CommonItem
}

func GetKey(userID int) (map[string]types.AttributeValue, error) {
	selectedKeys := map[string]string{
		"$pk": table.UserToPK(userID),
		"$sk": table.UserToSK(),
	}

	key, err := attributevalue.MarshalMap(selectedKeys)
	if err != nil {
		return nil, err
	}

	return key, nil
}
