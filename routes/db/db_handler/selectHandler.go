package db_handler

import (
	"TogetherAndStronger/routes/db/query"
	"encoding/json"
	"net/http"
)

func Select(w http.ResponseWriter, req *http.Request) {
	// Parse the request body into a map[string]interface{} object
	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the table name and columns from the request body
	tableName, ok := data["table"].(string)
	if !ok {
		http.Error(w, "missing or invalid 'table' parameter", http.StatusBadRequest)
		return
	}

	columnsInterface, ok := data["columns"].([]interface{})
	if !ok {
		http.Error(w, "missing or invalid 'columns' parameter", http.StatusBadRequest)
		return
	}

	var columns []string
	for _, col := range columnsInterface {
		colStr, ok := col.(string)
		if !ok {
			http.Error(w, "invalid 'columns' parameter", http.StatusBadRequest)
			return
		}
		columns = append(columns, colStr)
	}

	conditions, ok := data["conditions"].(map[string]interface{})
	if !ok {
		http.Error(w, "missing or invalid 'conditions' parameter", http.StatusBadRequest)
		return
	}

	// Call the SelectQuery function with the table name, columns and conditions
	rows, err := query.SelectQuery(tableName, columns, conditions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Get the column names from the sql.Rows object
	columnStrings, err := rows.Columns()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// INIT RESPONSE VAR DYNAMIC
	var response []map[string]interface{}
	for rows.Next() {
		rowData := make(map[string]interface{})
		scanArgs := make([]interface{}, len(columnStrings))
		for i := range columnStrings {
			scanArgs[i] = &scanArgs[i]
		}
		if err := rows.Scan(scanArgs...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i, value := range scanArgs {
			switch v := value.(type) {
			case []byte:
				rowData[columnStrings[i]] = string(v)
			default:
				rowData[columnStrings[i]] = v
			}
		}
		response = append(response, rowData)
	}

	// Check if any rows were returned
	if len(response) == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No rows found.",
		})
		return
	}

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
