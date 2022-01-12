package item

import (
	database "grpc-todo/database"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAllItems() ([]Item, error) {
	items := []Item{}

	scanner := database.CassandraClient.Query(
		`SELECT id, title, description, closed
		FROM ` + database.Keyspace + `.` + database.Table + `;`,
	).Iter().Scanner()

	for scanner.Next() {
		item := Item{}
		err := scanner.Scan(
			&item.ID, &item.Title, &item.Description, &item.Closed,
		)
		if err != nil {
			log.Println("error scanning cassandra response into struct : ", err)
			return nil, status.Error(codes.Internal, "something went wrong")
		}
		items = append(items, item)
	}
	err := scanner.Err()
	if err != nil {
		log.Println("error scanning item : ", err)
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	return items, nil
}
