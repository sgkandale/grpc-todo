package item

import (
	"context"
	database "grpc-todo/database"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) DeleteItem() error {

	_, err := database.ToDoItemsCollection.DeleteOne(
		context.TODO(),
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: item.ID,
			},
		},
	)

	if err != nil {
		log.Println("error item DeleteItem() : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
