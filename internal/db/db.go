package db

import (
	"errors"
	"log"

	"github.com/anil1226/go-gRPC/internal/rocket"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

type Database struct {
	Client *dynamodb.DynamoDB
}

const (
	containerEmp = "rockets"
)

func NewDatabase() (*Database, error) {

	// Create a new AWS session with shared credentials from ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),                         // Specify your AWS region
		Credentials: credentials.NewSharedCredentials("", "default"), // Use the default credentials profile
	})
	if err != nil {
		log.Fatal("Error creating session:", err)
	}

	// Create a DynamoDB client
	svc := dynamodb.New(sess)

	return &Database{
		Client: svc,
	}, nil
}

func (d *Database) GetRocketByID(id string) (rocket.Rocket, error) {
	result, err := d.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(containerEmp),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		return rocket.Rocket{}, nil
	}

	if result.Item == nil {
		msg := "Could not find '" + id + "'"
		return rocket.Rocket{}, errors.New(msg)
	}

	itemResponseBody := rocket.Rocket{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &itemResponseBody)
	if err != nil {
		log.Printf("Failed to unmarshal Record: %v", err)
		return rocket.Rocket{}, nil
	}
	return itemResponseBody, nil

}

func (d *Database) InsertRocket(rkt rocket.Rocket) error {
	rkt.ID = uuid.NewV4().String()

	av, err := dynamodbattribute.MarshalMap(rkt)
	if err != nil {
		log.Printf("Got error marshalling item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(containerEmp),
	}

	_, err = d.Client.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return err
	}

	return nil
}

func (d *Database) DeleteRocket(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(containerEmp),
	}

	_, err := d.Client.DeleteItem(input)
	if err != nil {
		log.Printf("Got error calling DeleteItem: %s", err)
		return err
	}
	return nil
}
