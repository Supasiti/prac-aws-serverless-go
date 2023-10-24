package user

import (
	"time"

	"github.com/supasiti/prac-aws-serverless-go/internal/store/table"
)

type IDGenerator = func() int

func CreateUser(params *CreateUserParams, idFn IDGenerator) *User {
	userID := idFn()

	return &User{
		FirstName: params.FirstName,
		Balance:   params.Balance,
		Email:     params.Email,
		UserID:    int(userID),
	}
}

func (u User) ToUserItem() Item {
	lastUpdated := time.Now()

	return Item{
		User: u,
		CommonItem: table.CommonItem{
			PK:          table.UserToPK(u.UserID),
			SK:          table.UserToSK(),
			Schema:      table.User,
			LastUpdated: lastUpdated,
			Created:     lastUpdated,
		},
	}
}
