package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func LoginSalary(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)
		fmt.Println(data["token"])

		selectQuery, err := query.SelectQuery("participant", []string{"idPARTICIPANT", "idCLIENT", "nom", "prenom"}, map[string]interface{}{"token": data["token"]})
		if err != nil {
			fmt.Println(w, err)
		}
		defer selectQuery.Close()

		if !selectQuery.Next() {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusNotFound)
			return
		} else {
			var idp, idc int
			var nom, prenom string
			err = selectQuery.Scan(&idp, &idc, &nom, &prenom)
			if err != nil {
				fmt.Println(err)
				return
			}
			libraries.Response(w, map[string]interface{}{
				"message": "Successfully logged in",
				"idp":     idp,
				"idc":     idc,
				"nom":     nom,
				"prenom":  prenom,
			}, http.StatusOK)
		}
	}
}
