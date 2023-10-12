package api

import "github.com/aws/aws-lambda-go/events"

type SuccessBody struct {
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *events.APIGatewayProxyResponse {
	body := &SuccessBody{
		Data: data,
	}
	res := NewApiResponseBuilder(body).Build()

	return res
}
