package database

import (
	"fmt"
	"log"
	"time"

	config "grpc-todo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var ToDoDB = config.Config.MySQL.DBName

const ItemsTable = "items"

func MySQLConn() *sqlx.DB {

	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?tls=false",
			config.Config.MySQL.Username,
			config.Config.MySQL.Password,
			config.Config.MySQL.Host,
			config.Config.MySQL.DBName,
		),
	)

	if err != nil {
		log.Fatal("error connecting to listings database : ", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Println("Connected to listings database")

	return db
}

var TodoDB = MySQLConn()
