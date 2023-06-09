package query

import (
	db2 "TogetherAndStronger/routes/db/init"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"strings"
)

// InsertQuery add a row to an existing table.
//
//	Example : db.InsertQuery("users", data := map[string]string{
//			"name":  "Johnny",
//			"email": "johnny@example.com",
//		})
func InsertQuery(tableName string, data map[string]interface{}) (int, error) {
	var columns []string
	var values []interface{}

	for col, val := range data {
		columns = append(columns, col)
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Trim(strings.Repeat("?,", len(columns)), ","))

	db, err := db2.InitDB()
	if err != nil {
		return 0, err
	}

	fmt.Println(html.EscapeString(query))
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer stmt.Close()

	fmt.Println(values)
	result, err := stmt.Exec(values...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Printf("%d row(s) affected\n", rowsAffected)
	return int(lastInsertId), nil
}
