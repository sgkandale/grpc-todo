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

func GetAllItems() ([]Item, error) {
	items := []Item{}

	result, err := database.TodoDB.Scan(
		&dynamodb.ScanInput{
			TableName: aws.String(database.TodoItemsTable),
		},
	)
	if err != nil {
		log.Println("error fetching all items from dynamodb : ", err)
		return items, status.Error(codes.Internal, "something went wrong")
	}

	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Println("error unmarshaling dynamodb response in struct : ", err)
			return items, status.Error(codes.Internal, "something went wrong")
		}

		items = append(items, item)
	}

	return items, nil
}
