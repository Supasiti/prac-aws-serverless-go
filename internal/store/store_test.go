package store

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	dbclient "github.com/supasiti/prac-aws-serverless-go/internal/dynamodb"
	"github.com/supasiti/prac-aws-serverless-go/internal/dynamodb/mocks"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user/stub"
)

func Test_store_GetUser(t *testing.T) {
	type fields struct {
		client func(t *testing.T) dbclient.DbClient
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				client: func(t *testing.T) dbclient.DbClient {
					m := mocks.NewDbClient(t)
					m.EXPECT().
						GetItem(mock.Anything, mock.Anything).
						Return(stub.UserGetItemOutput(), nil).Times(1)
					return m
				},
			},
			args:    args{userID: 1},
			want:    stub.User(),
			wantErr: assert.NoError,
		},
		{
			name: "should handle client error",
			fields: fields{
				client: func(t *testing.T) dbclient.DbClient {
					m := mocks.NewDbClient(t)
					m.EXPECT().
						GetItem(mock.Anything, mock.Anything).
						Return(nil, errors.New("some error")).Times(1)
					return m
				},
			},
			args:    args{userID: 1},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should user not found error",
			fields: fields{
				client: func(t *testing.T) dbclient.DbClient {
					m := mocks.NewDbClient(t)
					m.EXPECT().
						GetItem(mock.Anything, mock.Anything).
						Return(&dynamodb.GetItemOutput{}, nil).Times(1)
					return m
				},
			},
			args: args{userID: 1},
			want: nil,
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, user.ErrUserNotFound, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &store{
				client:    tt.fields.client(t),
				tableName: "mock-user-table",
			}
			got, err := s.GetUser(context.Background(), tt.args.userID)

			tt.wantErr(t, err, "store.GetUser(ctx, %d)", tt.args.userID)

			assert.Equalf(t, tt.want, got, "store.GetUser(ctx, %d) = %v , want %v", tt.args.userID, got, tt.want)
		})
	}
}

func Test_store_CreateUser(t *testing.T) {
	type fields struct {
		client func(t *testing.T) dbclient.DbClient
	}
	type args struct {
		params *user.CreateUserParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				client: func(t *testing.T) dbclient.DbClient {
					m := mocks.NewDbClient(t)
					m.EXPECT().
						PutItem(mock.Anything, mock.Anything).
						Return(&dynamodb.PutItemOutput{}, nil).Times(1)
					return m
				},
			},
			args:    args{params: stub.CreateUserParams()},
			want:    stub.User(),
			wantErr: assert.NoError,
		},
		{
			name: "should handle client error",
			fields: fields{
				client: func(t *testing.T) dbclient.DbClient {
					m := mocks.NewDbClient(t)
					m.EXPECT().
						PutItem(mock.Anything, mock.Anything).
						Return(nil, errors.New("some error")).Times(1)
					return m
				},
			},
			args:    args{params: stub.CreateUserParams()},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &store{
				client:    tt.fields.client(t),
				tableName: "mock-user-table",
				idGenerator: func() int {
					return 1
				},
			}
			got, err := s.CreateUser(context.Background(), tt.args.params)

			tt.wantErr(t, err, "store.CreateUser(ctx, %v)", tt.args.params)

			assert.Equalf(t, tt.want, got, "store.CreateUser(ctx, %v) = %v", tt.args.params, got)
		})
	}
}
