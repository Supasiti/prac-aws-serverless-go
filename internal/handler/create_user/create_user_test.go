package createuser

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/mocks"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user/stub"
	"github.com/supasiti/prac-aws-serverless-go/pkg/api"
)

func Test_handler_CreateUser(t *testing.T) {
	type fields struct {
		store func(*testing.T) store.Store
	}
	type args struct {
		body string
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
					m.EXPECT().CreateUser(mock.Anything, mock.Anything).
						Return(stub.User(), nil).Times(1)
					return m
				},
			},
			args: args{body: stub.CreateUserParams().String()},
			want: api.NewSuccessResponse(stub.User()),
		},
		{
			name: "should handle missing firstname",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					return m
				},
			},
			args: args{
				body: `{"balance": 3214.0, "email":"bob@email.com" }`,
			},
			want: api.NewValidationError(ErrMissingFirstName),
		},
		{
			name: "should handle missing email",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					return m
				},
			},
			args: args{
				body: `{"balance": 3214.0, "firstName":"bob" }`,
			},
			want: api.NewValidationError(ErrMissingEmail),
		},
		{
			name: "should handle incorrect json string",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					return m
				},
			},
			args: args{
				body: `{"balance": "3214.0", "firstName":"bob", "email": "bob@email.com" }`,
			},
			want: api.NewValidationError(ErrBadRequestBody),
		},
		{
			name: "should handle general error",
			fields: fields{
				store: func(t *testing.T) store.Store {
					m := mocks.NewStore(t)
					m.EXPECT().CreateUser(mock.Anything, mock.Anything).
						Return(nil, errors.New("some error")).Times(1)
					return m
				},
			},
			args: args{body: stub.CreateUserParams().String()},
			want: api.NewGeneralError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				store: tt.fields.store(t),
			}
			req := events.APIGatewayProxyRequest{
				Body: tt.args.body,
			}

			got := h.CreateUser(context.Background(), req)

			assert.Equalf(t, tt.want, got, "handler.CreateUser() = %+v, want %+v", got, tt.want)

		})
	}
}
