package item

import (
	"database/sql"
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) LoadDetails() error {

	err := database.TodoDB.Get(
		item,
		`SELECT * FROM `+database.ItemsTable+` 
		 WHERE id=?`,
		item.ID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return status.Error(codes.NotFound, "item not found")
		}
		log.Println("Error loading item from DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
