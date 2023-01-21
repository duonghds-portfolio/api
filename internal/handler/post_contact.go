package handler

import (
	"duonghds-portfolio-api/model"
	"duonghds-portfolio-api/service/dynamo"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

const (
	PortfolioContactTableName = "portfolio_contact"
)

func PostContact(body string) (events.APIGatewayProxyResponse, error) {
	data := PostContactData{}
	err := ParseData(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	msg, code, ok := ValidateData(&data)
	if !ok {
		fmt.Println("validate data failed: " + msg)
		return CreateResponse(msg, code)
	}
	newContactID, err := uuid.NewRandom()
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("post data is valid")

	model := model.PortfolioContactModel{
		UUID:        newContactID.String(),
		RealName:    data.RealName,
		Email:       data.Email,
		TextContent: data.TextContent,
		CreatedAt:   time.Now(),
	}
	err = dynamo.PutContactItem(model, PortfolioContactTableName)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("put item to dynamo success")
	return CreateResponse("success", http.StatusOK)
}

func ValidateData(data *PostContactData) (string, int, bool) {
	_, err := mail.ParseAddress(data.Email)
	if err != nil {
		return "error: invalid email format", http.StatusInternalServerError, false
	}
	return "", 0, true
}

type PostContactData struct {
	RealName    string `json:"realName"`
	Email       string `json:"email"`
	TextContent string `json:"textContent"`
}
