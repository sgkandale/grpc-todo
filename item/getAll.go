package item

import (
	"log"

	database "grpc-todo/database"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAllItems() ([]Item, error) {
	items := []Item{}

	rows, err := database.TodoDB.Queryx(
		`SELECT id, title, description, closed
		 FROM ` + database.ItemsTable + `
		 ORDER BY id DESC`,
	)

	if err != nil {
		log.Println("Error getting items from DB: ", err)
		return items, status.Error(codes.Internal, "something went wrong")
	}

	defer rows.Close()

	for rows.Next() {

		item := Item{}

		err := rows.StructScan(&item)

		if err != nil {
			log.Println("Error loading item from DB: ", err)
			return items, status.Error(codes.Internal, "something went wrong")
		}

		items = append(items, item)
	}

	return items, nil
}
