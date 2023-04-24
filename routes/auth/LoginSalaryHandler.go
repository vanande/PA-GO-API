package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"net/http"
)

func LoginSalary(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)
		println(data["token"])

		selectQuery, err := query.SelectQuery("participant", []string{"idPARTICIPANT", "idCLIENT", "idEQUIPE"}, map[string]interface{}{"token": data["token"]})
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
			var idp, idc, ide int
			err = selectQuery.Scan(&idp, &idc, &ide)
			if err != nil {
				return
			}
			println(idp, idc, ide)
			libraries.Response(w, map[string]interface{}{
				"message": "Successfully logged in",
				"idp":     idp,
				"idc":     idc,
				"ide":     ide,
			}, http.StatusNotFound)
		}
	}
}
