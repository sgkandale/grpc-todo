package primaryDB

import (
	"context"
	"fmt"
	"log"

	config "grpc-todo/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ToDoDB = config.Config.Postgres.DBName

const ItemsTable = "items"

func NewPostgresConn() *pgxpool.Pool {

	log.Println("connecting Postgres")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Config.Postgres.Username, config.Config.Postgres.Password,
		config.Config.Postgres.Endpoint, config.Config.Postgres.Port, ToDoDB,
	)

	conn, err := pgxpool.Connect(
		context.TODO(),
		connStr,
	)
	if err != nil {
		log.Fatal("error opening Postgres : ", err)
	}

	pingErr := conn.Ping(context.TODO())
	if pingErr != nil {
		log.Fatal("error pinging Postgres : ", pingErr)
	}

	log.Println("Postgres Connected")

	return conn
}

var TodoDB = NewPostgresConn()
