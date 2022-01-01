package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	config "grpc-todo/config"
	server "grpc-todo/server"
	todoPB "grpc-todo/todoPB"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func getServerTLS() credentials.TransportCredentials {
	creds, err := credentials.NewServerTLSFromFile(
		config.Config.ServerConfig.CertPath,
		config.Config.ServerConfig.KeyPath,
	)
	if err != nil {
		log.Fatal("error in server TLS cert : ", err)
	}

	return creds
}

func main() {

	log.Println("starting...")

	serverConfig := config.Config.ServerConfig

	lis, err := net.Listen("tcp", "0.0.0.0:"+serverConfig.Port)
	if err != nil {
		log.Fatal(err)
	}

	var grpcServer *grpc.Server
	tlsStat := ""

	if serverConfig.TLSEnabled {
		tlsStat = "secure"
		grpcServer = grpc.NewServer(
			grpc.Creds(getServerTLS()),
		)
	} else {
		tlsStat = "insecure"
		grpcServer = grpc.NewServer()
	}

	todoPB.RegisterTodoServiceServer(grpcServer, &server.Server{})

	go func() {
		log.Println(tlsStat, "gRPC Server Started on :"+serverConfig.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// Stop the server
	log.Println("stopping the server")
	grpcServer.Stop()
	log.Println("server stopped")
}
