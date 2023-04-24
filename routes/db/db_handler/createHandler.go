package db_handler

import (
	"TogetherAndStronger/routes/db/query"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, req *http.Request) {
	// Parse the request body into a map[string]interface{} object
	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the table name and data from the request body
	tableName, ok := data["table"].(string)
	if !ok {
		http.Error(w, "missing or invalid 'table' parameter", http.StatusBadRequest)
		return
	}

	createData, ok := data["data"].(map[string]interface{})
	if !ok {
		http.Error(w, "missing or invalid 'data' parameter", http.StatusBadRequest)
		return
	}

	// Call the CreateQuery function with the table name and create data
	_, err := query.InsertQuery(tableName, createData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a JSON response with a success message and data
	response := map[string]interface{}{
		"message": "successfully created",
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
