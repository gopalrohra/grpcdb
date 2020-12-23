package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/gopalrohra/grpc_db/grpc_database"
	_ "github.com/lib/pq"

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
	log.Printf("Received: %v", r.GetDbname())
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=postgres"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return &pb.DatabaseResult{Status: "Error"}, nil
	}
	defer db.Close()
	query := fmt.Sprintf("create database %v", r.GetDbname())
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return &pb.DatabaseResult{Status: "Error"}, nil
	}
	return &pb.DatabaseResult{Status: "Success"}, nil
}
