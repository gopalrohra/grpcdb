package database

import (
	"fmt"
	"strings"

	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type PostgresDBQueryBuilder struct{}

func (qb PostgresDBQueryBuilder) GetDSN(d *pb.DatabaseInfo, selectDB bool) string {

	if selectDB {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.GetHost(), d.GetPort(), d.GetUser(), d.GetPassword(), d.GetName())
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", d.GetHost(), d.GetPort(), d.GetUser(), d.GetPassword())

}

func (qb PostgresDBQueryBuilder) GetDBCreationQuery(d *pb.DatabaseInfo) string {

	return fmt.Sprintf("create database %v", d.GetName())
}

func (qb PostgresDBQueryBuilder) GetTableCreationQuery(tableRequest *pb.TableRequest) string {

	return fmt.Sprintf("create table %s (%s)", tableRequest.GetName(), strings.Join(tableRequest.GetColumnDef(), ","))
}

func (qb PostgresDBQueryBuilder) GetSelectionQuery(sq *pb.SelectQuery) string {

	fields := strings.Join(sq.GetFields(), ",")
	tableName := sq.GetTableName()
	clauses := strings.Join(sq.GetClauses(), " and ")
	groupBy := strings.Join(sq.GetGroupby(), ", ")
	orderBy := strings.Join(sq.GetOrderby(), ", ")
	if clauses != "" {
		clauses = " where " + clauses
	}
	if groupBy != "" {
		groupBy = " group by " + groupBy
	}
	if orderBy != "" {
		orderBy = " order by " + orderBy
	}
	return fmt.Sprintf("select %s from %s %s %s %s", fields, tableName, clauses, groupBy, orderBy)
}

func (qb PostgresDBQueryBuilder) GetInsertionQuery(iq *pb.InsertQueryRequest) string {

	return fmt.Sprintf("insert into %s(%s)values(%s)", iq.GetTableName(), strings.Join(iq.GetColumns(), ","), strings.Join(iq.GetColumnValues(), ","))
}

func (qb PostgresDBQueryBuilder) GetUpdationQuery(updateQuery *pb.UpdateQuery) string {

	columnValues := getColumnValues(updateQuery.Columns)
	return fmt.Sprintf("update %s set %s where %s", updateQuery.GetTableName(), columnValues, strings.Join(updateQuery.GetClauses(), " and "))
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
