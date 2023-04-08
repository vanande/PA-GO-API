package db

import (
	db2 "TogetherAndStronger/routes/db/init"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func InsertQuery(tableName string, data map[string]string) {
	var columns []string
	var values []interface{}

	for col, val := range data {
		columns = append(columns, col)
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Trim(strings.Repeat("?,", len(columns)), ","))

	db, err := db2.InitDB()
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d row(s) affected\n", rowsAffected)
}
