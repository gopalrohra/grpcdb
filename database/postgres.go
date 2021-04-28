package database

import (
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

// Postgres struct to implement Database interface
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

func fetchRows(query string, psqlInfo string) (*sql.Rows, error) {
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
func (p Postgres) CreateDatabase(d *pb.DatabaseInfo) (*pb.DatabaseResult, error) {
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=postgres", d.GetHost(), d.GetUser(), d.GetPassword())
	query := fmt.Sprintf("create database %v", d.GetName())
	_, err := executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return &pb.DatabaseResult{Status: "Error", Description: err.Error()}, nil
	}
	return &pb.DatabaseResult{Status: "Success", Description: "Database created."}, nil
}

// CreateTable method to create a new table
// and returns a TableResponse with Status either "Success" or "Error"
func (p Postgres) CreateTable(t *pb.TableRequest) (*pb.TableResponse, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", t.Info.GetHost(), t.Info.GetPort(), t.Info.GetUser(), t.Info.GetPassword(), t.Info.GetName())
	query := fmt.Sprintf("create table %s (%s)", t.GetName(), strings.Join(t.GetColumnDef(), ","))
	fmt.Println(query)
	_, err := executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	return &pb.TableResponse{Status: "Success", Description: "Table created"}, nil
}

// ExecuteSelect methods creates a select query
// and returns the result
func (p Postgres) ExecuteSelect(sq *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", sq.Info.GetHost(), sq.Info.GetPort(), sq.Info.GetUser(), sq.Info.GetPassword(), sq.Info.GetName())
	fields := strings.Join(sq.GetFields(), ",")
	tableName := sq.GetTableName()
	clauses := strings.Join(sq.GetClauses(), " and ")
	if clauses != "" {
		clauses = " where " + clauses
	}
	query := fmt.Sprintf("select %s from %s %s", fields, tableName, clauses)
	fmt.Println(query)
	result, err := fetchRows(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	defer result.Close()
	rows, err := getRows(result)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	// fmt.Println(rows)
	//fmt.Println(fmt.Sprintf(" %d Rows: %v", len(rows), rows[0]))
	return &pb.SelectQueryResult{Status: "Success", Description: "Query fetched some result", Records: rows}, nil
}
func getRows(result *sql.Rows) ([]*pb.Row, error) {
	defer result.Close()
	columns, err := result.Columns()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Columns: %v\n", columns)
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

// ExecuteInsert inserts  a record in a given table
func (p Postgres) ExecuteInsert(iq *pb.InsertQueryRequest) (*pb.InsertQueryResponse, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", iq.Info.GetHost(), iq.Info.GetPort(), iq.Info.GetUser(), iq.Info.GetPassword(), iq.Info.GetName())
	query := fmt.Sprintf("insert into %s(%s)values(%s)", iq.GetTableName(), strings.Join(iq.GetColumns(), ","), strings.Join(iq.GetColumnValues(), ","))
	result, err := executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}

	return &pb.InsertQueryResponse{Status: "Success", Description: "Record inserted", InsertedId: fmt.Sprintf("%v", lastInsertID)}, nil
}

// ExecuteUpdate updates a record in a given table
func (p Postgres) ExecuteUpdate(updateQuery *pb.UpdateQuery) (*pb.UpdateQueryResult, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", updateQuery.Info.GetHost(), updateQuery.Info.GetPort(), updateQuery.Info.GetUser(), updateQuery.Info.GetPassword(), updateQuery.Info.GetName())
	columnValues := getColumnValues(updateQuery.Columns)
	query := fmt.Sprintf("update %s set %s where %s", updateQuery.GetTableName(), columnValues, strings.Join(updateQuery.GetClauses(), " and "))
	fmt.Println(query)
	result, err := executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	return &pb.UpdateQueryResult{Status: "Success", Description: "Record updated", RowsAffected: fmt.Sprintf("%v", rowsAffected)}, nil
}
func getColumnValues(columns []*pb.Column) string {
	var resultColumns []string
	for _, col := range columns {
		resultColumns = append(resultColumns, col.GetColumnName()+"="+col.GetColumnValue())
	}
	if len(resultColumns) == 0 {
		fmt.Println("Error as no columns specified")
		return ""
	}
	return strings.Join(resultColumns, ", ")
}
