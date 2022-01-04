package item

import (
	"log"

	database "grpc-todo/database"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) SetClosed() error {

	_, err := database.TodoDB.UpdateItem(
		&dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":c": {
					BOOL: aws.Bool(true),
				},
			},
			TableName: aws.String(database.TodoItemsTable),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(item.ID),
				},
			},
			ReturnValues:     aws.String("UPDATED_NEW"),
			UpdateExpression: aws.String("set closed = :c"),
		},
	)

	if err != nil {
		log.Println("Error closing item from DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
