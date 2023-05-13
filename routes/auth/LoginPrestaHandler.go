package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

<<<<<<< HEAD
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func LoginPresta(w http.ResponseWriter, req *http.Request) {

=======
func LoginPresta(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
	switch req.Method {
	case "POST":

		enableCors(&w)

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
<<<<<<< HEAD
=======
			println(id)
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c

			role := "prestataire"
			token, err := libraries.CreateToken(role, id)
			if err != nil {
<<<<<<< HEAD
				fmt.Errorf("Error while generating the token : %v", err)
=======
				fmt.Errorf("Error while generating token: %v", err)
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			}

			libraries.Response(w, map[string]interface{}{
				"message": "Successfully logged in",
				"id":      id,
				"token":   token,
			}, http.StatusOK)
		}

	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	}
}
