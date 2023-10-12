package user

import (
	"time"

	"github.com/supasiti/prac-aws-serverless-go/internal/store/table"
)

type IdGenerator = func() int

func CreateUser(params *CreateUserParams, idFn IdGenerator) *User {
	userID := idFn()

	return &User{
		FirstName: params.FirstName,
		Balance:   params.Balance,
		Email:     params.Email,
		UserID:    int(userID),
	}
}

func (u User) ToUserItem() UserItem {
	lastUpdated := time.Now()

	return UserItem{
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
