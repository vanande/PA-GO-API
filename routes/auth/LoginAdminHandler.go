package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
	"strings"
)

func LoginAdmin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		nom, OK := data["nom"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusUnauthorized)
			return
		}

		password, OK := data["password"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusUnauthorized)
			return
		}

		selectQuery, err := query.SelectQuery("admin", []string{"idADMIN", "role"}, map[string]interface{}{"nom": strings.ToLower(nom), "password": password})
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
			if rang == 1 {
				role = "employee"
			} else if rang == 2 {
				role = "admin"
			} else if rang == 3 {
				role = "superadmin"
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
