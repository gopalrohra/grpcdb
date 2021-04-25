package database

import (
	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

// Database is an interface which defines operations that can be performed.
type Database interface {
	CreateDatabase(*pb.DatabaseInfo) (*pb.DatabaseResult, error)
	CreateTable(*pb.TableRequest) (*pb.TableResponse, error)
	ExecuteSelect(*pb.SelectQuery) (*pb.SelectQueryResult, error)
	ExecuteInsert(*pb.InsertQueryRequest) (*pb.InsertQueryResponse, error)
}
