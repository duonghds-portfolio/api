package dynamo

import (
	"duonghds-portfolio-api/model"
	"fmt"
	"os"

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
	result, err := Client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("list table dynamo error")
		return err
	}
	for _, n := range result.TableNames {
		fmt.Println("table: ", *n)
	}
	_, err = Client.PutItem(input)
	if err != nil {
		fmt.Println("put item to dynamo error")
		return err
	}
	fmt.Println()
	return nil
}
