package query

import (
	db2 "TogetherAndStronger/routes/db/init"
	"database/sql"
	"fmt"
	"strings"
)

type name struct {
	nom string
	age int
}

// SelectQuery returns a sql.Rows object. It takes a table name, a slice of columns and a map of conditions.
//
// Example : SelectQuery("users", []string{"id", "name"}, map[string]interface{}{"id": 1})
func SelectQuery(tableName string, columns []string, conditions map[string]interface{}) (*sql.Rows, error) {
	// Building base query : SELECT column1, column2, ... FROM table
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), tableName)

	// Magic begins here
	var values []interface{}
	if len(conditions) > 0 {
		var where []string
		for column, value := range conditions {
			where = append(where, fmt.Sprintf("%s = ?", column))
			values = append(values, value)
		}
		query += " WHERE " + strings.Join(where, " AND ")
	}

	db, err := db2.InitDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
