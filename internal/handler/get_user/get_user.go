package getuser

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
	"github.com/supasiti/prac-aws-serverless-go/pkg/api"
)

var (
	ErrMissingUserID = errors.New("missing userID")
	ErrIncorrectType = errors.New("userID must be an integer")
)

type GetUserHandler interface {
	GetUser(context.Context, events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse
}

type getUserHandler struct {
	store store.Store
}

func NewHandler(s store.Store) GetUserHandler {
	handler := getUserHandler{
		store: s,
	}
	return &handler
}

func (h *getUserHandler) GetUser(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	log.Printf("GetUser Event: %+v", req)

	userID, err := validateRequest(req)
	if err != nil {
		return api.NewValidationError(err)
	}

	result, err := h.store.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return api.NewDataNotFoundError(err)
		}
		return api.NewGeneralError()
	}

	return api.NewSuccessResponse(result)
}

func validateRequest(req events.APIGatewayProxyRequest) (int, error) {
	param := req.PathParameters["userID"]
	if len(param) == 0 {
		return 0, ErrMissingUserID
	}

	userID, err := strconv.Atoi(param)
	if err != nil {
		return 0, ErrIncorrectType
	}

	return userID, nil
}
