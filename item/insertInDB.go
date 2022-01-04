package item

import (
	"log"

	database "grpc-todo/database"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) InsertInDB() error {

	_, err := database.TodoDB.NamedExec(
		`INSERT INTO `+database.ItemsTable+`
		 (id, title, description, closed)
		 values(:id, :title, :description, :closed)`,
		map[string]interface{}{
			"id":          item.ID,
			"title":       item.Title,
			"description": item.Description,
			"closed":      item.Closed,
		},
	)

	if err != nil {
		log.Println("Error inserting item in DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
