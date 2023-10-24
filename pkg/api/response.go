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

type ResponseBuilder interface {
	Build() *events.APIGatewayProxyResponse
}

type Builder struct {
	resp events.APIGatewayProxyResponse
}

func NewBuilder(body interface{}) *Builder {

	resp := events.APIGatewayProxyResponse{
		Headers:    defaultHeader,
		StatusCode: http.StatusOK,
	}
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)

	return &Builder{resp: resp}
}

func (r *Builder) WithStatus(status int) *Builder {
	r.resp.StatusCode = status
	return r
}

func (r *Builder) WithHeaders(headers map[string]string) *Builder {
	for k, v := range headers {
		r.resp.Headers[k] = v
	}
	return r
}

func (r *Builder) Build() *events.APIGatewayProxyResponse {
	return &r.resp
}
