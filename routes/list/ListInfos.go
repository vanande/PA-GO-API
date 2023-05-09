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

		// Magic start here
		var dataRows []map[string]interface{} // Create a slice of hashmap to store the data
		for selectQuery.Next() {
			columnNames, _ := selectQuery.Columns()                 // Get the column names to init the map
			columnValues := make([]interface{}, len(columnNames))   // Create a slice(= improved array -> very opti) of interface{} to store the values
			valuePointers := make([]interface{}, len(columnValues)) // Create a slice of interface{} to store the pointers of the multiple &var, &var2, &var3, etc.
			for i := range columnValues {
				valuePointers[i] = &columnValues[i] // Store the pointers of the multiple &var, &var2, &var3, etc. in the slice
			}
			err = selectQuery.Scan(valuePointers...) // Scan need addresses of the variables to store the values
			if err != nil {
				fmt.Println(err)
				continue
			}
			// INIT DONE
			// Now we have the column names and the values stored in the slice columnValues
			// We can now create a map with the column names as keys and the values as values like this:
			// columnNames = ["id", "name", "description"]
			// columnValues = ["1", "Super duper test", "Super description"]

			// This part is just to create the map with the column names as keys and the values as values
			// dataRow = {"id": "1", "name": "Super duper test", "description": "Super description"}
			dataRow := make(map[string]interface{})
			for i, val := range columnValues {
				if val != nil {
					switch v := val.(type) {
					case []byte:
						dataRow[columnNames[i]] = string(v) // Small hacks to avoid strange values like é$ù&é"àç_è-ç�
					default:
						dataRow[columnNames[i]] = v // Example : dataRow["id"] = "1"
					}
				}
			}
			dataRows = append(dataRows, dataRow) // Obviously append the dataRow to the slice of map
			// dataRows = [{"id": "1", "name": "Super duper test", "description": "Super description"}]
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
