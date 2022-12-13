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
	port                = ":3099"
	invalidDatabaseType = "Not valid database type"
)

var databases = map[string]db.Database{
	"postgres": db.Database{DriverName: "postgres", QBuilder: new(db.GenericSQLQueryBuilder), DSNBuilder: new(db.PostgresDSNBuilder)},
}

type server struct {
	pb.UnimplementedGRPCDatabaseServer
}

func databaseImplementations() map[string]db.Database {

	return databases
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
func (s *server) CreateDatabase(ctx context.Context, r *pb.DatabaseInfo) (*pb.DatabaseResult, error) {
	log.Printf("Received: %v", r)

	dbObj, isPresent := databaseImplementations()[r.Type]
	if !isPresent {
		return &pb.DatabaseResult{Status: "Error", Description: invalidDatabaseType}, nil
	}
	result, err := dbObj.CreateDatabase(r)
	if err != nil {
		return &pb.DatabaseResult{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
func (s *server) CreateTable(ctx context.Context, r *pb.TableRequest) (*pb.TableResponse, error) {
	log.Printf("Received: %v", r)

	dbObj, isPresent := databaseImplementations()[r.Info.Type]
	if !isPresent {
		return &pb.TableResponse{Status: "Error", Description: invalidDatabaseType}, nil
	}
	result, err := dbObj.CreateTable(r)
	if err != nil {
		return &pb.TableResponse{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
func (s *server) ExecuteSelect(ctx context.Context, r *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	log.Printf("Received: %v", r)

	dbObj, isPresent := databaseImplementations()[r.Info.Type]
	if !isPresent {
		return &pb.SelectQueryResult{Status: "Error", Description: invalidDatabaseType}, nil
	}
	result, err := dbObj.ExecuteSelect(r)
	if err != nil {
		return &pb.SelectQueryResult{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}

func (s *server) ExecuteInsert(ctx context.Context, r *pb.InsertQueryRequest) (*pb.InsertQueryResponse, error) {
	log.Printf("Received: %v", r)

	dbObj, isPresent := databaseImplementations()[r.Info.Type]
	if !isPresent {
		return &pb.InsertQueryResponse{Status: "Error", Description: invalidDatabaseType}, nil
	}
	result, err := dbObj.ExecuteInsert(r)
	if err != nil {
		return &pb.InsertQueryResponse{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
func (s *server) ExecuteUpdate(ctx context.Context, r *pb.UpdateQuery) (*pb.UpdateQueryResult, error) {
	log.Printf("Received: %v\n", r)

	dbObj, isPresent := databaseImplementations()[r.Info.Type]
	if !isPresent {
		return &pb.UpdateQueryResult{Status: "Error", Description: invalidDatabaseType}, nil
	}
	result, err := dbObj.ExecuteUpdate(r)
	if err != nil {
		return &pb.UpdateQueryResult{Status: "Error", Description: err.Error()}, nil
	}
	return result, nil
}
