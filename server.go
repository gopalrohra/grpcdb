package main

import (
	"context"
	"fmt"
	"log"
	"net"

	db "github.com/gopalrohra/grpcdb/database"
	pb "github.com/gopalrohra/grpcdb/grpc_database"
	_ "github.com/lib/pq"
	grpc "google.golang.org/grpc"
)

const (
	port = ":3099"
)

type server struct {
	pb.UnimplementedGRPCDatabaseServer
}

func databaseImplementations() map[string]db.Database {
	return map[string]db.Database{"postgres": db.Postgres{}}
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

	result, err := databaseImplementations()["postgres"].CreateDatabase(r)
	if err != nil {
		return &pb.DatabaseResult{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
func (s *server) CreateTable(ctx context.Context, r *pb.TableRequest) (*pb.TableResponse, error) {
	log.Printf("Received: %v", r)

	result, err := databaseImplementations()["postgres"].CreateTable(r)
	if err != nil {
		return &pb.TableResponse{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
func (s *server) ExecuteSelect(ctx context.Context, r *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	log.Printf("Received: %v", r)

	result, err := databaseImplementations()["postgres"].ExecuteSelect(r)
	if err != nil {
		return &pb.SelectQueryResult{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
