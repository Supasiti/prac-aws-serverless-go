package getuser

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/mocks"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user/stub"
	"github.com/supasiti/prac-aws-serverless-go/pkg/api"
)

func Test_getUserHandler_GetUser(t *testing.T) {
	type fields struct {
		store func(t *testing.T) store.Store
	}
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *events.APIGatewayProxyResponse
	}{
		{
			name: "success",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					m.EXPECT().
						GetUser(mock.Anything, 1).
						Return(stub.GetUser(), nil).Times(1)
					return m
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{
						"userID": "1",
					},
				},
			},
			want: api.NewSuccessResponse(stub.GetUser()),
		},
		{
			name: "should handle missing userID",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					return m
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{},
				},
			},
			want: api.NewValidationError(ErrMissingUserID),
		},
		{
			name: "should handle incorrect userID type",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					return m
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{
						"userID": "invalid id",
					},
				},
			},
			want: api.NewValidationError(ErrIncorrectType),
		},
		{
			name: "should handle user not found error",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					m.EXPECT().
						GetUser(mock.Anything, 1).
						Return(nil, user.ErrUserNotFound).
						Times(1)
					return m
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{
						"userID": "1",
					},
				},
			},
			want: api.NewDataNotFoundError(user.ErrUserNotFound),
		},
		{
			name: "should handle user other error",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					m.EXPECT().
						GetUser(mock.Anything, 1).
						Return(nil, errors.New("some error")).
						Times(1)
					return m
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{
						"userID": "1",
					},
				},
			},
			want: api.NewGeneralError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &getUserHandler{
				store: tt.fields.store(t),
			}
			got := h.GetUser(context.Background(), tt.args.req)

			assert.Equalf(t, tt.want, got, "getUserHandler.GetUser() = %v, want %v", got, tt.want)

		})
	}
}
