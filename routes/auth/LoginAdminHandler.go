package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
)

func LoginAdmin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

<<<<<<< HEAD
		selectQuery, err := query.SelectQuery("admin", []string{"idADMIN", "role"}, map[string]interface{}{"nom": data["nom"], "password": data["password"]})
=======
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
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
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
<<<<<<< HEAD
			if rang < 2 {
				role = "employee"
			} else {
				role = "admin"
=======
			if rang < 1 {
				role = "noob"
			} else if rang < 2 {
				role = "admin"
			} else {
				role = "superadmin"
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
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
