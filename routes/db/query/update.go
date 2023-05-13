package query

import (
	db2 "TogetherAndStronger/routes/db/init"
	"fmt"
	"strings"
)

// UpdateQuery updates rows in an existing table based on specified conditions.
//
// Example usage:
//
//	updates := map[string]interface{}{
//			"token": 2222,
//	}
//
//	conditions := map[string]interface{}{
//			"id": 1,
//	}
//
// err := UpdateQuery("participant", updates, conditions)
//
//	if err != nil {
//			fmt.Println(err)
//			return
//	}
func UpdateQuery(tableName string, updates map[string]interface{}, conditions map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	var set []string
	var values []interface{}
	for column, value := range updates {
		set = append(set, fmt.Sprintf("%s = ?", column))
		values = append(values, value)
	}

	query := fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(set, ", "))

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
