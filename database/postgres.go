package database

import (
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type Postgres struct{}

func executeQuery(query string, psqlInfo string) (sql.Result, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (p Postgres) CreateDatabase(d *pb.Database) (*pb.DatabaseResult, error) {
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=postgres"
	query := fmt.Sprintf("create database %v", d.GetDbname())
	_, err := executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return &pb.DatabaseResult{Status: "Error", Description: err.Error()}, nil
	}
	return &pb.DatabaseResult{Status: "Success", Description: "Database created."}, nil
}

func (p Postgres) CreateTable(t *pb.TableRequest) (*pb.TableResponse, error) {
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres password=postgres dbname=%s", t.Info.GetDbname())
	query := fmt.Sprintf("create table %s (%s)", t.GetName(), strings.Join(t.GetColumnDef(), ","))
	fmt.Println(query)
	_, err := executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return nil, err
	}
	return &pb.TableResponse{Status: "Success", Description: "Table created"}, nil
}
func (p Postgres) ExecuteSelect(sq *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	return &pb.SelectQueryResult{}, nil
}
