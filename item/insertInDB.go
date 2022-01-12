package item

import (
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) InsertInDB() error {

	err := database.CassandraClient.Query(
		`INSERT INTO `+database.Table+`
		(id, title, description, closed)
		VALUES (?, ?, ?, ?)`,
		item.ID, item.Title, item.Description, item.Closed,
	).Exec()
	if err != nil {
		log.Println("Cassandra insertion error : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
