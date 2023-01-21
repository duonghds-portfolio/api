package main

import (
	"duonghds-portfolio-api/internal/handler"
	"duonghds-portfolio-api/service/dynamo"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	PostContactPath = "/api/contact"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod == http.MethodGet {
		return events.APIGatewayProxyResponse{
			Body:       "error: method get not allowed with this url",
			StatusCode: http.StatusMethodNotAllowed,
		}, nil
	}
	switch req.Path {
	case PostContactPath:
		dynamo.Init()
		return handler.PostContact(req.Body)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "error: url/api not exists",
			StatusCode: http.StatusNotFound,
			Headers: map[string]string{
				"Access-Control-Allow-Headers": "Content-Type",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "*",
			},
		}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
	// HandleRequest(events.APIGatewayProxyRequest{
	// 	HTTPMethod: http.MethodPost,
	// 	Path:       LoginPath,
	// 	Body:       "{ \"username\": \"duong7\", \"password\": \"password_duong\" }",
	// })
}
