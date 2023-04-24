package auth

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"net/http"
)

func LoginCompany(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)
		println(data["email"])

		selectQuery, err := query.SelectQuery("client", []string{"idCLIENT"}, map[string]interface{}{"email": data["email"]})
		if err != nil {
			return
		}
		defer selectQuery.Close()

		if !selectQuery.Next() {
			libraries.Response(w, "Invalid data", http.StatusNotFound)
			return
		} else {
			var id int
			err = selectQuery.Scan(&id)
			if err != nil {
				return
			}
			println(id)
			libraries.Response(w, "Successfully logged in", http.StatusOK)
		}
	}
}
