package database

import (
	"database/sql"
	"fmt"

	pb "github.com/gopalrohra/grpc_db/grpc_database"
	_ "github.com/lib/pq"
)

type Postgres struct{}

func (p Postgres) CreateDatabase(d *pb.Database) (*pb.DatabaseResult, error) {
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=postgres"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return &pb.DatabaseResult{Status: "Error"}, nil
	}
	defer db.Close()
	query := fmt.Sprintf("create database %v", d.GetDbname())
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return &pb.DatabaseResult{Status: "Error"}, nil
	}
	return &pb.DatabaseResult{Status: "Success"}, nil

}

func (p Postgres) CreateTable(t *pb.TableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}
func (p Postgres) ExecuteSelect(sq *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	return &pb.SelectQueryResult{}, nil
}
