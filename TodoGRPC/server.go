package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"todo/todo"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 14586))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := todo.Server{}
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}
