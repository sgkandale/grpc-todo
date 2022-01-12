package item

import (
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) DeleteItem() error {

	err := database.CassandraClient.Query(
		`DELETE FROM `+database.Keyspace+`.`+database.Table+`
		WHERE id = ?`,
		item.ID,
	).Exec()
	if err != nil {
		log.Println("Cassandra deletion error : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
