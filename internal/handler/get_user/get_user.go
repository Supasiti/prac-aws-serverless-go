package getuser

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/supasiti/prac-aws-serverless-go/internal/pkg/json"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
	"github.com/supasiti/prac-aws-serverless-go/pkg/api"
)

var (
	ErrMissingUserID = errors.New("missing userID")
	ErrIncorrectType = errors.New("userID must be an integer")
)

type Handler interface {
	GetUser(context.Context, events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse
}

type handler struct {
	store store.Store
}

func NewHandler(s store.Store) Handler {
	h := handler{
		store: s,
	}
	return &h
}

func (h *handler) GetUser(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	log.Printf("GetUser Event: %s", json.ToInlineJSON(req))

	userID, err := validateRequest(req)
	if err != nil {
		log.Printf("GetUser error: %+v", err)
		return api.NewValidationError(err)
	}

	result, err := h.store.GetUser(ctx, userID)
	if err != nil {
		log.Printf("GetUser error: %+v", err)

		if errors.Is(err, user.ErrUserNotFound) {
			return api.NewDataNotFoundError(err)
		}
		return api.NewGeneralError()
	}

	log.Printf("GetUser result: %s", json.ToInlineJSON(result))
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
