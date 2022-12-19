package database

import (
	"fmt"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type MySQLDSNBuilder struct{}

func (builder *MySQLDSNBuilder) GetDSN(d *pb.DatabaseInfo, selectDB bool) string {

	if selectDB {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.GetUser(), d.GetPassword(), d.GetHost(), d.GetPort(), d.GetName())
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", d.GetUser(), d.GetPassword(), d.GetHost(), d.GetPort())

}
