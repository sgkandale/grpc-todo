package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) LoadDetails() error {

	row := database.TodoDB.QueryRow(
		context.Background(),
		`SELECT title, description, closed
		 FROM `+database.ItemsTable+`
		 WHERE id=$1`,
		item.ID,
	)

	err := row.Scan(&item.Title, &item.Description, &item.Closed)

	if err != nil {
		if err == pgx.ErrNoRows {
			return status.Error(codes.NotFound, "item not found")
		}
		log.Println("Error loading item from DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
