package config

import (
	config "grpc-todo/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var TodoItemsTable = config.Config.DynamoDB.Table

func TodoDBConn() *dynamodb.DynamoDB {

	mySession := session.Must(
		session.NewSession(
			&aws.Config{
				Region: aws.String(config.Config.DynamoDB.Region),
				Credentials: credentials.NewStaticCredentials(
					config.Config.DynamoDB.AccessKey,
					config.Config.DynamoDB.SecretKey,
					"",
				),
				// Endpoint: aws.String("http://localhost:8000"), // for local use
			},
		),
	)

	client := dynamodb.New(mySession)

	return client
}

var TodoDB = TodoDBConn()
