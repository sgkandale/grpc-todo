package item

import (
	"log"

	database "grpc-todo/database"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) DeleteItem() error {

	_, err := database.TodoDB.NamedExec(
		`DELETE FROM `+database.ItemsTable+`
		 WHERE id=:id`,
		map[string]interface{}{
			"id": item.ID,
		},
	)

	if err != nil {
		log.Println("Error deleting item from DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
