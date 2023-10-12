package stub

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
)

var (
	itemOutput = map[string]types.AttributeValue{
		"$created":     &types.AttributeValueMemberS{Value: "2023/02/13 12:21:20"},
		"$lastUpdated": &types.AttributeValueMemberS{Value: "2023/02/13 12:21:20"},
		"$pk":          &types.AttributeValueMemberS{Value: "1"},
		"$schema":      &types.AttributeValueMemberS{Value: "USER"},
		"$sk":          &types.AttributeValueMemberS{Value: "USER"},
		"balance":      &types.AttributeValueMemberN{Value: "3251.12"},
		"email":        &types.AttributeValueMemberS{Value: "bob.builder@email.com"},
		"firstName":    &types.AttributeValueMemberS{Value: "Bob"},
		"userID":       &types.AttributeValueMemberN{Value: "1"},
	}
)

func User() *user.User {
	return &user.User{
		UserID:    1,
		Balance:   3251.12,
		FirstName: "Bob",
		Email:     "bob.builder@email.com",
	}
}

func UserGetItemOutput() *dynamodb.GetItemOutput {
	return &dynamodb.GetItemOutput{Item: itemOutput}
}

func CreateUserParams() *user.CreateUserParams {
	return &user.CreateUserParams{
		Balance:   3251.12,
		FirstName: "Bob",
		Email:     "bob.builder@email.com",
	}
}
