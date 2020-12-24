package main

import (
	"context"
	"fmt"
	"log"
	"net"

	db "github.com/gopalrohra/grpc_db/database"
	pb "github.com/gopalrohra/grpc_db/grpc_database"
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
	fmt.Println("Opening the port 3099")
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
func (s *server) CreateDatabase(ctx context.Context, r *pb.Database) (*pb.DatabaseResult, error) {
	log.Printf("Received: %v", r)
	supportedDatabases := map[string]db.Database{"postgres": db.Postgres{}}
	result, err := supportedDatabases["postgres"].CreateDatabase(r)
	if err != nil {
		return &pb.DatabaseResult{Status: "Error"}, nil
	}
	return result, nil
}
