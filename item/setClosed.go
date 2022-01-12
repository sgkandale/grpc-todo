package item

import (
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) SetClosed() error {

	err := database.CassandraClient.Query(
		`UPDATE `+database.Table+`
		SET closed = ?
		WHERE id = ?`,
		true, item.ID,
	).Exec()
	if err != nil {
		log.Println("Cassandra update error : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
