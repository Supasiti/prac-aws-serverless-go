package api

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/supasiti/prac-aws-serverless-go/pkg/http/header"
)

var (
	defaultHeader = map[string]string{
		header.ContentType:  "application/json",
		header.CacheControl: "no-cache",
	}
)

type ApiResponseBuilder interface {
	Build() *events.APIGatewayProxyResponse
}

type apiResponseBuilder struct {
	resp events.APIGatewayProxyResponse
}

func NewApiResponseBuilder(body interface{}) *apiResponseBuilder {

	resp := events.APIGatewayProxyResponse{
		Headers:    defaultHeader,
		StatusCode: http.StatusOK,
	}
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)

	return &apiResponseBuilder{resp: resp}
}

func (r *apiResponseBuilder) WithStatus(status int) *apiResponseBuilder {
	r.resp.StatusCode = status
	return r
}

func (r *apiResponseBuilder) WithHeaders(headers map[string]string) *apiResponseBuilder {
	for k, v := range headers {
		r.resp.Headers[k] = v
	}
	return r
}

func (r *apiResponseBuilder) Build() *events.APIGatewayProxyResponse {
	return &r.resp
}
