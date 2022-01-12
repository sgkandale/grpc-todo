package item

import (
	"context"
	database "grpc-todo/database"
	"log"
	"time"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) LoadDetails() error {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := database.CassandraClient.Query(`
		SELECT title, description, closed
		FROM `+database.Keyspace+`.`+database.Table+`
		WHERE id = ? LIMIT 1`,
		item.ID,
	).WithContext(ctx).Scan(&item.Title, &item.Description, &item.Closed)
	if err != nil {
		if err == gocql.ErrNotFound {
			return status.Error(codes.NotFound, "item not found")
		}
		log.Println("Error finding item in Cassandra : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
