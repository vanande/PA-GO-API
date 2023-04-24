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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
