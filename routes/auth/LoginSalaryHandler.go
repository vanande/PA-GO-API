package auth

import (
	"encoding/json"
	"net/http"
)

func LoginSalary(w http.ResponseWriter, req *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
