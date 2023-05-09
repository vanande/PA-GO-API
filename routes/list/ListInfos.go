package list

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func ListInfos(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		selectQuery, err := query.SelectQuery("infos", []string{"*"}, map[string]interface{}{})
		if err != nil {
			fmt.Println(w, err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}
		defer selectQuery.Close()

		dataRows := []map[string]interface{}{}
		for selectQuery.Next() {
			columnNames, _ := selectQuery.Columns()
			columnValues := make([]interface{}, len(columnNames))
			valuePointers := make([]interface{}, len(columnValues))
			for i := range columnValues {
				valuePointers[i] = &columnValues[i]
			}
			err = selectQuery.Scan(valuePointers...)
			if err != nil {
				fmt.Println(err)
				continue
			}

			dataRow := make(map[string]interface{})
			for i, val := range columnValues {
				if val != nil {
					switch v := val.(type) {
					case []byte:
						dataRow[columnNames[i]] = string(v)
					default:
						dataRow[columnNames[i]] = v
					}
				}
			}
			dataRows = append(dataRows, dataRow)
		}

		if len(dataRows) == 0 {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusNotFound)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    dataRows,
		}, http.StatusOK)
	}
}
