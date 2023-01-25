package dynamo

import (
	"duonghds-portfolio-api/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var Client *dynamodb.DynamoDB

func Init() {
	fmt.Println("start connect dynamodb, region: " + os.Getenv("COMMON_REGION"))
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("COMMON_REGION"))},
	)
	if err != nil {
		fmt.Printf("got error when create session config: %s\n", err.Error())
		return
	}
	// Create DynamoDB client
	Client = dynamodb.New(sess)
	fmt.Println("connected dynamodb")
}

func PutContactItem(item model.PortfolioContactModel, tableName string) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("marshal item to dynamo attribute error")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = Client.PutItem(input)
	if err != nil {
		fmt.Println("put item to dynamo error")
		return err
	}
	return nil
}

func GetNotes(tableName string) (*[]model.Note, error) {
	output, err := Client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		fmt.Println("scan notes table in dynamodb failed: ", err.Error())
		return nil, err
	}
	notes := []model.Note{}
	for _, item := range output.Items {
		note := model.Note{}
		err = dynamodbattribute.UnmarshalMap(item, &note)

		if err != nil {
			fmt.Println("got error unmarshalling: ", err.Error())
			return nil, err
		}
		notes = append(notes, note)
	}
	return &notes, nil
}

func PutNote(item model.Note, tableName string) error {
	noteTTLStr := os.Getenv("NOTE_TTL")
	noteTTL, _ := strconv.Atoi(noteTTLStr)
	item.ExpiredTime.Add(time.Second * time.Duration(noteTTL))
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("marshal item to dynamo attribute error")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = Client.PutItem(input)
	if err != nil {
		fmt.Println("put item to dynamo error")
		return err
	}
	return nil
}
