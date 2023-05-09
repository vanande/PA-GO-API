package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func LoginPresta(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)
		println(data["email"])

		selectQuery, err := query.SelectQuery("prestataire", []string{"idPRESTATAIRE"}, map[string]interface{}{"email": data["email"]})
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
			var id int
			err = selectQuery.Scan(&id)
			if err != nil {
				return
			}
			println(id)

			role := "prestataire"
			token, err := libraries.CreateToken(role, id)
			if err != nil {
				fmt.Errorf("Error while generating token: %v", err)
			}

			libraries.Response(w, map[string]interface{}{
				"message": "Successfully logged in",
				"id":      id,
				"token":   token,
			}, http.StatusOK)
		}
	}
}
