package query

import (
	db "TogetherAndStronger/routes/db/init"
	"fmt"
)

// DeleteQuery deletes a row from a table.
//
// Example : DeleteQuery("users", "id = ?", 1)
func DeleteQuery(tableName string, condition string, values ...interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, condition)

	db, err := db.InitDB()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("%d row(s) affected\n", rowsAffected)
	return nil
}
