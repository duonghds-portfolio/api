package handler

import (
	"duonghds-portfolio-api/model"
	"duonghds-portfolio-api/service/dynamo"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

const (
	PortfolioNoteTableName = "portfolio_note"
)

func GetNotes(body string) (events.APIGatewayProxyResponse, error) {
	notes, err := dynamo.GetNotes(PortfolioNoteTableName)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	notesByte, err := json.Marshal(notes)
	if err != nil {
		fmt.Println("marshal notes to json failed")
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return CreateResponse(string(notesByte), http.StatusOK)
}

func PutNote(body string) (events.APIGatewayProxyResponse, error) {
	data := NoteData{}
	err := ParseData(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}

	if data.NoteID > 14 || data.NoteID < 0 {
		return CreateResponse("note_id is not exist", http.StatusNotFound)
	}
	if data.ReadOnly {
		return CreateResponse("this note cannot override", http.StatusConflict)
	}
	if data.ExpiredTime.After(time.Now()) {
		return CreateResponse("this note is expired yet", http.StatusConflict)
	}
	if data.Text == "" {
		return CreateResponse("cannot put blank text", http.StatusConflict)
	}
	fmt.Printf("note data: %+v", data)
	item := model.Note{
		NoteID:      data.NoteID,
		Text:        data.Text,
		ReadOnly:    data.ReadOnly,
		ExpiredTime: data.ExpiredTime,
	}
	err = dynamo.PutNote(item, PortfolioNoteTableName)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return CreateResponse("success", http.StatusOK)
}

type NoteData struct {
	NoteID      int       `json:"noteID"`
	Text        string    `json:"text"`
	ReadOnly    bool      `json:"readOnly"`
	ExpiredTime time.Time `json:"expiredTime"`
}
