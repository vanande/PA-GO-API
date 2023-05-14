package libraries

import (
	"encoding/json"
	"net/http"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Body(w http.ResponseWriter, req *http.Request) map[string]interface{} {
	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return data
}
