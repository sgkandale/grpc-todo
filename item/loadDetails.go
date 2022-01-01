package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) LoadDetails() error {

	res := database.ToDoItemsCollection.FindOne(
		context.TODO(),
		bson.M{"_id": item.ID},
	)

	err := res.Decode(item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return status.Error(codes.NotFound, "item not found")
		}
		log.Println("error item LoadDetails() : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil

}
