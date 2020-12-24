package database

import (
	pb "github.com/gopalrohra/grpc_db/grpc_database"
)

type Database interface {
	CreateDatabase(*pb.Database) (*pb.DatabaseResult, error)
	CreateTable(*pb.TableRequest) (*pb.TableResponse, error)
	ExecuteSelect(*pb.SelectQuery) (*pb.SelectQueryResult, error)
}
