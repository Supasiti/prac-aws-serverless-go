package table

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrSchemaNotFound = errors.New("schema not found")
)

type CommonItem struct {
	PK          string    `dynamodbav:"$pk"`
	SK          string    `dynamodbav:"$sk"`
	Schema      Schema    `dynamodbav:"$schema"`
	LastUpdated time.Time `dynamodbav:"$lastUpdated"`
	Created     time.Time `dynamodbav:"$created"`
}

func UserToPK(userID int) string {
	return fmt.Sprintf("%d", userID)
}

func UserToSK() string {
	return "USER"
}
