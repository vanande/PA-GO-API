package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"net/http"
	"fmt"
)

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func LoginPresta(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got TWOOOOOOOOOOO")

	switch req.Method {
		case "POST":

			// enableCors(&w)

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
				}, http.StatusNotFound)
				return
			} else {
				var id int
				err = selectQuery.Scan(&id)
				if err != nil {
					return
				}
				println(id)
				libraries.Response(w, map[string]interface{}{
					"message": "Successfully logged in",
					"id":      id,
				}, http.StatusOK)
			}


		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		}
}
