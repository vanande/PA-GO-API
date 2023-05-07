package libraries

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, message map[string]interface{}, code int) {

	response := map[string]interface{}{}
	for key, value := range message {
		response[key] = value
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
