package database

import (
	"database/sql"
	"fmt"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type Database struct {
	DriverName string
	QBuilder   QueryBuilder
	DSNBuilder DataSourceNameBuilder
}

// Database is an interface which defines operations that can be performed.
type IDatabase interface {
	CreateDatabase(*pb.DatabaseInfo) (*pb.DatabaseResult, error)
	CreateTable(*pb.TableRequest) (*pb.TableResponse, error)
	ExecuteSelect(*pb.SelectQuery) (*pb.SelectQueryResult, error)
	ExecuteInsert(*pb.InsertQueryRequest) (*pb.InsertQueryResponse, error)
	ExecuteUpdate(*pb.UpdateQuery) (*pb.UpdateQueryResult, error)
}

func (database Database) executeQuery(query string, psqlInfo string) (sql.Result, error) {
	db, err := sql.Open(database.DriverName, psqlInfo)
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

func (database Database) executeInsertQuery(query string, psqlInfo string, returnValueExpected bool) (string, error) {
	db, dbErr := sql.Open(database.DriverName, psqlInfo)
	if dbErr != nil {
		return "", dbErr
	}
	defer db.Close()

	returnID := ""

	var err error
	if returnValueExpected {
		//result, err = db.Exec(query).scan(&returnID)
		err = db.QueryRow(query).Scan((&returnID))
	} else {
		_, err = db.Exec(query)
	}
	if err != nil {
		return "", err
	}
	return returnID, nil
}

func (database Database) fetchRows(query string, psqlInfo string) (*sql.Rows, error) {
	db, err := sql.Open(database.DriverName, psqlInfo)
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

// CreateDatabase method creates a new database
// and returns DatabaseResult defined in grpcdb package
func (database *Database) CreateDatabase(d *pb.DatabaseInfo) (*pb.DatabaseResult, error) {

	psqlInfo := database.DSNBuilder.GetDSN(d, false)
	query := database.QBuilder.GetDBCreationQuery(d)
	_, err := database.executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return &pb.DatabaseResult{Status: "Error", Description: err.Error()}, nil
	}
	return &pb.DatabaseResult{Status: "Success", Description: "Database created."}, nil
}

// CreateTable method to create a new table
// and returns a TableResponse with Status either "Success" or "Error"
func (database *Database) CreateTable(t *pb.TableRequest) (*pb.TableResponse, error) {
	psqlInfo := database.DSNBuilder.GetDSN(t.Info, true)
	query := database.QBuilder.GetTableCreationQuery(t)
	fmt.Println(query)
	_, err := database.executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	return &pb.TableResponse{Status: "Success", Description: "Table created"}, nil
}

// ExecuteSelect methods creates a select query
// and returns the result
func (database *Database) ExecuteSelect(sq *pb.SelectQuery) (*pb.SelectQueryResult, error) {
	psqlInfo := database.DSNBuilder.GetDSN(sq.Info, true)
	query := database.QBuilder.GetSelectionQuery(sq)
	fmt.Println(query)
	result, err := database.fetchRows(query, psqlInfo)
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
func (database *Database) ExecuteInsert(iq *pb.InsertQueryRequest) (*pb.InsertQueryResponse, error) {
	psqlInfo := database.DSNBuilder.GetDSN(iq.Info, true)
	query := database.QBuilder.GetInsertionQuery(iq)
	fmt.Println("Query to be executed", query)
	if iq.ReturningIdColumnName != "" {
		query = fmt.Sprintf("%s returning %s", query, iq.ReturningIdColumnName)
	}
	result, err := database.executeInsertQuery(query, psqlInfo, (iq.ReturningIdColumnName != ""))
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	// lastInsertID, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Printf("Error occured: %v\n", err)
	// 	return nil, err
	// }

	return &pb.InsertQueryResponse{Status: "Success", Description: "Record inserted", InsertedId: fmt.Sprintf("%v", result)}, nil
}

// ExecuteUpdate updates a record in a given table
func (database *Database) ExecuteUpdate(updateQuery *pb.UpdateQuery) (*pb.UpdateQueryResult, error) {
	psqlInfo := database.DSNBuilder.GetDSN(updateQuery.Info, true)
	query := database.QBuilder.GetUpdationQuery(updateQuery)
	fmt.Println(query)
	result, err := database.executeQuery(query, psqlInfo)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	return &pb.UpdateQueryResult{Status: "Success", Description: "Record updated", RowsAffected: fmt.Sprintf("%v", rowsAffected)}, nil
}
