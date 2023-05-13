package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func AddActivityOption(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		//											-->	change here <--
		_, err := query.InsertQuery("tas.option", map[string]interface{}{
			// Data like that
			"idActivite":      data["id"],
			"idlist_activite": data["ida"],
			"idlist_option":   data["ido"],
		})
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully added",
		}, http.StatusOK)
	}
}
