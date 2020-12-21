package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/gopalrohra/grpc-db/grpc_database"

	grpc "google.golang.org/grpc"
)

const (
	port = ":3099"
)

type server struct {
	pb.UnimplementedGRPCDatabaseServer
}

func main() {
	fmt.Println("Starting the grpc database service.")
	fmt.Printf("Opening the port")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterGRPCDatabaseServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
