package database

import (
	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type QueryBuilder interface {
	GetDBCreationQuery(*pb.DatabaseInfo) string
	GetTableCreationQuery(*pb.TableRequest) string
	GetSelectionQuery(*pb.SelectQuery) string
	GetInsertionQuery(*pb.InsertQueryRequest) string
	GetUpdationQuery(*pb.UpdateQuery) string
}