package database

import (
	"fmt"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type PostgresDSNBuilder struct{}

func (builder PostgresDSNBuilder) GetDSN(d *pb.DatabaseInfo, selectDB bool) string {

	if selectDB {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.GetHost(), d.GetPort(), d.GetUser(), d.GetPassword(), d.GetName())
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", d.GetHost(), d.GetPort(), d.GetUser(), d.GetPassword())

}
