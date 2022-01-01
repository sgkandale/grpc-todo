package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) InsertInDB() error {

	_, err := database.ToDoItemsCollection.InsertOne(
		context.TODO(),
		item,
	)

	if err != nil {
		log.Println("error item InsertInDB() : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
