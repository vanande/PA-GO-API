package query

import (
	"database/sql"
	"fmt"
	"strings"

	db2 "TogetherAndStronger/routes/db/init"
)

//	SelectWithInnerJoin returns a sql.Rows object by executing an inner join query on two tables.
//	It takes a slice of table names, a slice of columns, a slice of maps containing the join conditions for two consecutive tables,
//	and a map containing the WHERE conditions for the query.
//
// Example usage:
//
//	cols := []string{"users.name", "orders.order_id", "orders.price"}
//	tables := []string{"users", "orders"}
//	joins := []map[string]interface{}{
//	"users.user_id": "orders.user_id"},
//	}
//	conditions := map[string]interface{}{
//		"orders.price >": 50,
//	}
//	rows, err := SelectWithInnerJoin(tables, cols, joins, conditions)
//
//	if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer rows.Close()
//
//		var res []struct {
//			Name string
//			OrderID int
//			Price int
//		}
//
//		for rows.Next() {
//			var r struct {
//				Name string
//				OrderID int
//				Price int
//			}
//				if err := rows.Scan(&r.Name, &r.OrderID, &r.Price); err != nil {
//				fmt.Println(err)
//				return
//			}
//			res = append(res, r)
//		}
//
//		fmt.Println(res)
func SelectWithInnerJoin(tables []string, columns []string, joins []map[string]string, conditions map[string]interface{}) (*sql.Rows, error) {

	if len(tables) < 2 {
		return nil, fmt.Errorf("tables : %d < 2. Need at least 2", len(tables))
	}
	if len(tables)-len(joins) != 1 {
		return nil, fmt.Errorf("tables : %d, joins : %d. Need n+1 tables than there is joins", len(tables), len(joins))
	}

	// base query : SELECT column1, column2, ... FROM table1
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), tables[0])

	// join tables
	for i, join := range joins {
		table := tables[i+1]
		query += fmt.Sprintf(" INNER JOIN %s ON ", table)

		var joinConditions []string
		for id1, id2 := range join {
			joinConditions = append(joinConditions, fmt.Sprintf("%s = %s", id1, id2))
		}
		query += strings.Join(joinConditions, " AND ")
	}

	// WHERE condition
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
