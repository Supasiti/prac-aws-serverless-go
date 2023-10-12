package api

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrGeneralServer = errors.New("internal server error")
)

type ErrorBody struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) *apiResponseBuilder {
	errorBody := &ErrorBody{
		Message: err.Error(),
	}

	builder := NewApiResponseBuilder(errorBody).WithStatus(http.StatusInternalServerError)
	return builder
}

func NewGeneralError() *events.APIGatewayProxyResponse {
	return NewErrorResponse(ErrGeneralServer).Build()
}

func NewDataNotFoundError(err error) *events.APIGatewayProxyResponse {
	return NewErrorResponse(err).WithStatus(http.StatusNotFound).Build()
}

func NewValidationError(err error) *events.APIGatewayProxyResponse {
	return NewErrorResponse(err).WithStatus(http.StatusBadRequest).Build()
}
