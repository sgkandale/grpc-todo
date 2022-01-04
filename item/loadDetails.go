package item

import (
	database "grpc-todo/database"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) LoadDetails() error {

	result, err := database.TodoDB.GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String(database.TodoItemsTable),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(item.ID),
				},
			},
		},
	)
	if err != nil {
		log.Println("error fetching item from dynamodb : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	if result.Item == nil {
		return status.Error(codes.NotFound, "item not found")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		log.Println("error unmarshalling dynamodb response : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
