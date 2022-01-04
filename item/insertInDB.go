package item

import (
	"log"

	database "grpc-todo/database"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) InsertInDB() error {

	marshalledItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("error marshalling item for DynamoDB : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: aws.String(database.TodoItemsTable),
	}

	_, err = database.TodoDB.PutItem(input)

	if err != nil {
		log.Println("error putting item in DynamoDB : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
