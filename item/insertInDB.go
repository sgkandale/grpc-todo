package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) InsertInDB() error {

	_, err := database.TodoDB.Exec(
		context.Background(),
		`INSERT INTO `+database.ItemsTable+`
		 (id, title, description, closed)
		 values($1, $2, $3, $4)`,
		item.ID, item.Title, item.Description, item.Closed,
	)

	if err != nil {
		log.Println("Error inserting item in DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
