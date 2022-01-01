package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAllItems() ([]Item, error) {

	res, err := database.ToDoItemsCollection.Find(
		context.TODO(),
		bson.D{},
		options.Find().SetSort(
			bson.D{
				primitive.E{
					Key:   "_id",
					Value: -1,
				},
			},
		),
	)
	if err != nil {
		log.Println("error item GetAllItems() : ", err)
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	var items []Item
	err = res.All(context.TODO(), &items)
	if err != nil {
		log.Println("error item GetAllItems() : ", err)
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	return items, nil

}
