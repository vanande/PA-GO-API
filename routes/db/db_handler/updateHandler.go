package db_handler

import (
	"TogetherAndStronger/routes/db/query"
	"encoding/json"
	"net/http"
)

func Update(w http.ResponseWriter, req *http.Request) {
	// Parse the request body into a map[string]interface{} object
	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the table name, filter, and update data from the request body
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

	updateData, ok := data["data"].(map[string]interface{})
	if !ok {
		http.Error(w, "missing or invalid 'data' parameter", http.StatusBadRequest)
		return
	}

	// Call the UpdateQuery function with the table name, filter, and update data
	err := query.UpdateQuery(tableName, filter, updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a success response with a JSON message
	response := make(map[string]string)
	response["message"] = "Successfully updated the record(s)."
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
