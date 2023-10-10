package api

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorBody struct {
	Message string
}

func NewErrorResponse(err error) *apiResponseBuilder {
	errorBody := &ErrorBody{
		Message: err.Error(),
	}

	builder := NewApiResponseBuilder(errorBody).WithStatus(http.StatusInternalServerError)
	return builder
}

func NewGeneralError(err error) *events.APIGatewayProxyResponse {
	return NewErrorResponse(err).Build()
}

func NewDataNotFoundError(err error) *events.APIGatewayProxyResponse {
	return NewErrorResponse(err).WithStatus(http.StatusNotFound).Build()
}

func NewValidationError(err error) *events.APIGatewayProxyResponse {
	return NewErrorResponse(err).WithStatus(http.StatusBadRequest).Build()
}
