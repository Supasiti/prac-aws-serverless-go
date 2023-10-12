package createuser

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	jsonpkg "github.com/supasiti/prac-aws-serverless-go/internal/pkg/json"
	"github.com/supasiti/prac-aws-serverless-go/internal/store"
	"github.com/supasiti/prac-aws-serverless-go/internal/store/user"
	"github.com/supasiti/prac-aws-serverless-go/pkg/api"
)

var (
	ErrMissingFirstName = errors.New("missing firstName")
	ErrMissingEmail     = errors.New("missing email")
	ErrBadRequestBody   = errors.New("incorrect request body")
)

type Handler interface {
	CreateUser(context.Context, events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse
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

func (h *handler) CreateUser(ctx context.Context, req events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	log.Printf("CreateUser Event: %s", jsonpkg.ToInlineJSON(req))

	params, err := validateRequest(req)
	if err != nil {
		log.Printf("CreateUser error: %+v", err)
		return api.NewValidationError(err)
	}

	result, err := h.store.CreateUser(ctx, params)
	if err != nil {
		log.Printf("CreateUser error: %+v", err)
		return api.NewGeneralError()
	}

	log.Printf("CreateUser result: %s", jsonpkg.ToInlineJSON(result))
	return api.NewSuccessResponse(result)
}

func validateRequest(req events.APIGatewayProxyRequest) (*user.CreateUserParams, error) {
	var p user.CreateUserParams

	err := json.Unmarshal([]byte(req.Body), &p)
	if err != nil {
		return nil, ErrBadRequestBody
	}

	if len(p.FirstName) == 0 {
		return nil, ErrMissingFirstName
	}

	if len(p.Email) == 0 {
		return nil, ErrMissingEmail
	}

	// allow balance to be 0 if not specified
	return &p, nil
}
