package list

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ListActivities(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		enableCors(&w)
		selectQuery, err := query.SelectQuery("list_activite", []string{"*"}, map[string]interface{}{})
		if err != nil {
			fmt.Println(w, err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}
		defer selectQuery.Close()

		var rows []map[string]interface{}
		for selectQuery.Next() {
			columns, _ := selectQuery.Columns()
			values := make([]interface{}, len(columns))
			scanArgs := make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
			}
			err = selectQuery.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err)
				continue
			}

			row := make(map[string]interface{})
			for i, col := range values {
				if col != nil {
					switch v := col.(type) {
					case []byte:
						row[columns[i]] = string(v)
					default:
						row[columns[i]] = v
					}
				}
			}
			rows = append(rows, row)
		}

		if len(rows) == 0 {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusNotFound)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    rows,
		}, http.StatusOK)

	case "OPTIONS":
		fmt.Println("Preflight handled")
		w.WriteHeader(http.StatusOK)
	}
}
