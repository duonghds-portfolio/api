package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/valyala/fastjson"
)

func CreateResponse(msg string, code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: code,
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
		},
	}, nil
}

func ParseData(body string, model interface{}) error {
	err := fastjson.Validate(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &model)
	if err != nil {
		return err
	}
	return nil
}
