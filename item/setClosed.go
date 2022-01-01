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

func (item *Item) SetClosed() error {

	_, err := database.ToDoItemsCollection.UpdateOne(
		context.TODO(),
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: item.ID,
			},
		},
		bson.D{
			primitive.E{
				Key: "$set",
				Value: bson.D{
					primitive.E{
						Key:   "closed",
						Value: true,
					},
				},
			},
		},
	)

	if err != nil {
		log.Println("error item SetClosed() : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
