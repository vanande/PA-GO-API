package api

import (
	"TogetherAndStronger/routes/db/query"
	"encoding/json"
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	// Parse the request body into a map[string]interface{} object
	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the table name and filter from the request body
	tableName, ok := data["table"].(string)
	if !ok {
		http.Error(w, "missing or invalid 'table' parameter", http.StatusBadRequest)
		return
	}

	filter, ok := data["filter"].(map[string]interface{})
	if !ok {
		http.Error(w, "missing or invalid 'filter' parameter", http.StatusBadRequest)
		return
	}

	// Convert the filter map into a string representation
	var filterStr string
	for k, v := range filter {
		filterStr += fmt.Sprintf("%s='%v' AND ", k, v)
	}
	filterStr = filterStr[:len(filterStr)-5] // remove the trailing "AND "

	// Call the DeleteQuery function with the table name and filter
	err := query.DeleteQuery(tableName, filterStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a JSON response with a success message and data
	response := map[string]interface{}{
		"message": "successfully deleted",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
