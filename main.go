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
	GetNotesPath    = "/api/notes"
	PutNotePath     = "/api/note"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dynamo.Init()
	return handler.PostContact(req.Body)
	if req.HTTPMethod == http.MethodGet {
		switch req.Path {
		case GetNotesPath:
			dynamo.Init()
			return handler.GetNotes(req.Body)
		default:
			return handler.CreateResponse("error: url/api not exists", http.StatusNotFound)
		}
	}
	switch req.Path {
	case PostContactPath:
		dynamo.Init()
		return handler.PostContact(req.Body)
	case PutNotePath:
		dynamo.Init()
		return handler.PutNote(req.Body)
	default:
		return handler.CreateResponse("error: url/api not exists", http.StatusNotFound)
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
