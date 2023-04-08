package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// BuildInsertQuery builds an INSERT query from a map of data. It returns the query and the values to be inserted.
func BuildInsertQuery(tableName string, data map[string]string) (q string, v []string) {
	var columns []string
	var values []string
	var placeholders []string

	for col, val := range data {
		columns = append(columns, col)
		values = append(values, val)
		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	return query, values
}

// ExecInsertQuery executes the INSERT query built from the given table name and data.
// It returns the number of rows affected and an error if any.
func ExecInsertQuery(db *sql.DB, tableName string, data map[string]string) (int64, error) {
	query, stringValues := BuildInsertQuery(tableName, data)

	values := make([]interface{}, len(stringValues))
	for i, v := range stringValues {
		values[i] = v
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
