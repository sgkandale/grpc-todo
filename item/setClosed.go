package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) SetClosed() error {

	_, err := database.TodoDB.Exec(
		context.Background(),
		`UPDATE `+database.ItemsTable+`
		 SET closed=true
		 WHERE id=$1`,
		item.ID,
	)

	if err != nil {
		log.Println("Error closing item in DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
