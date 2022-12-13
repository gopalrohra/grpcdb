package database

import (
	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type DataSourceNameBuilder interface {
	GetDSN(*pb.DatabaseInfo, bool) string
}
