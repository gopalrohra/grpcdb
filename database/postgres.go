package database

import (
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

// Postgres struct to implement Database interface
type Postgres struct{}

func executeDDLQuery(query string, psqlInfo string) (sql.Result, error) {
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

func executeDMLQuery(query string, psqlInfo string) (*sql.Rows, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query(query)
	if err != nil {
		return result, err
	}
	return result, nil
}

//CreateDatabase method creates a new database
// and returns DatabaseResult defined in grpcdb package
func (p Postgres) CreateDatabase(d *pb.Database) (*pb.DatabaseResult, error) {
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=postgres"
	query := fmt.Sprintf("create database %v", d.GetDbname())
	_, err := executeDDLQuery(query, psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return &pb.DatabaseResult{Status: "Error", Description: err.Error()}, nil
	}
	return &pb.DatabaseResult{Status: "Success", Description: "Database created."}, nil
}

// CreateTable method to create a new table
// and returns a TableResponse with Status either "Success" or "Error"
func (p Postgres) CreateTable(t *pb.TableRequest) (*pb.TableResponse, error) {
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres password=postgres dbname=%s", t.Info.GetDbname())
	query := fmt.Sprintf("create table %s (%s)", t.GetName(), strings.Join(t.GetColumnDef(), ","))
	fmt.Println(query)
	_, err := executeDDLQuery(query, psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return nil, err
	}
	return &pb.TableResponse{Status: "Success", Description: "Table created"}, nil
}

// ExecuteSelect methods creates a select query
// and returns the result
func (p Postgres) ExecuteSelect(sq *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres password=postgres dbname=%s", sq.Info.GetDbname())
	query := fmt.Sprintf("select %s from %s", strings.Join(sq.GetFields(), ","), sq.GetTableName())
	fmt.Println(query)
	result, err := executeDMLQuery(query, psqlInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occured: %v", err))
		return nil, err
	}
	defer result.Close()
	rows, err := getRows(result)
	fmt.Println(rows)
	//fmt.Println(fmt.Sprintf(" %d Rows: %v", len(rows), rows[0]))
	return &pb.SelectQueryResult{Status: "Success", Description: "Query fetched some result", Records: rows}, nil
}
func getRows(result *sql.Rows) ([]*pb.Row, error) {
	defer result.Close()
	columns, err := result.Columns()
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Columns: %v", columns))
	rows := make([]*pb.Row, 0)
	for result.Next() {
		columnValues := make([]interface{}, len(columns))
		for i := range columnValues {
			columnValues[i] = new(interface{})
		}
		if err := result.Scan(columnValues...); err != nil {
			return nil, err
		}
		protoColumns := make([]*pb.Column, 0)
		for i, column := range columns {
			protoColumns = append(protoColumns, &pb.Column{ColumnName: column, ColumnValue: fmt.Sprintf("%v", *columnValues[i].(*interface{}))})
		}
		rows = append(rows, &pb.Row{Columns: protoColumns})
	}
	return rows, nil
}
