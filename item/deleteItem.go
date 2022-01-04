package item

import (
	"log"

	database "grpc-todo/database"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) DeleteItem() error {

	_, err := database.TodoDB.DeleteItem(
		&dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(item.ID),
				},
			},
			TableName: aws.String(database.TodoItemsTable),
		},
	)

	if err != nil {
		log.Println("Error deleting item from DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
