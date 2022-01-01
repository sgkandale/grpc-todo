package config

import (
	"log"

	"grpc-todo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func MongoDBClientConn() *mongo.Client {

	log.Println("connecting mongodb")

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(config.Config.MongoDBURI),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB Connected")

	return client
}


var MongoDBClient = MongoDBClientConn()

var ToDoDB = MongoDBClient.Database("grpcTodo")
var ToDoItemsCollection = ToDoDB.Collection("items")
