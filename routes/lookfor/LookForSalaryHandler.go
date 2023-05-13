package lookfor

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func LookForSalary(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

<<<<<<< HEAD
		selectQuery, err := query.SelectQuery("participant", []string{"*"}, map[string]interface{}{"idPARTICIPANT": data["id"], "idCLIENT": data["idc"]})
=======
		selectQuery, err := query.SelectQuery("participant", []string{"*"}, map[string]interface{}{"idPARTICIPANT": data["id"], "idEQUIPE": data["ide"], "idCLIENT": data["idc"]})
>>>>>>> 1784635def712a3f39b341ecc4c2c2f38ab4d056
		if err != nil {
			fmt.Println(w, err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}
		defer selectQuery.Close()

		if !selectQuery.Next() {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusNotFound)
			return
		} else {
			columns, _ := selectQuery.Columns()
			values := make([]interface{}, len(columns))
			scanArgs := make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
				fmt.Println("scanargs ", i, " ", scanArgs[i])
			}
			err = selectQuery.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err)
				return
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
					fmt.Printf("%s: %v\n", columns[i], row[columns[i]])
				}
			}

			fmt.Println("row=", row)

			libraries.Response(w, map[string]interface{}{
				"message": "Successfully fetched data",
				"data":    row,
			}, http.StatusOK)
		}
	}
}
