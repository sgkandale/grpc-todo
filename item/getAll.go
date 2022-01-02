package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAllItems() ([]Item, error) {
	items := []Item{}

	rows, err := database.TodoDB.Query(
		context.Background(),
		`SELECT id, title, description, closed
		 FROM `+database.ItemsTable+`
		 ORDER BY id DESC`,
	)

	if err != nil {
		log.Println("Error getting items from DB: ", err)
		return items, status.Error(codes.Internal, "something went wrong")
	}

	defer rows.Close()

	for rows.Next() {

		item := Item{}

		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Closed)

		if err != nil {
			log.Println("Error loading item from DB: ", err)
			return items, status.Error(codes.Internal, "something went wrong")
		}

		items = append(items, item)
	}

	return items, nil
}
