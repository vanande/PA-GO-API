package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func LoginAdmin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		selectQuery, err := query.SelectQuery("admin", []string{"idADMIN", "role"}, map[string]interface{}{"nom": data["nom"], "password": data["password"]})
		if err != nil {
			return
		}
		defer selectQuery.Close()

		if !selectQuery.Next() {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusUnauthorized)
			return
		} else {
			var id, rang int
			err = selectQuery.Scan(&id, &rang)
			if err != nil {
				return
			}
			println(id)

			var role string
			if rang < 2 {
				role = "employee"
			} else {
				role = "admin"
			}

			token, err := libraries.CreateToken(role, id)
			if err != nil {
				fmt.Errorf("error while generating token: %v", err)
			}

			libraries.Response(w, map[string]interface{}{
				"message": "Successfully logged in",
				"id":      id,
				"token":   token,
			}, http.StatusOK)
		}
	}
}
