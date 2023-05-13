package option

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func Make(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		idlo, OK := data["idlo"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID1",
			}, http.StatusBadRequest)
			return
		}

		idla, OK := data["idla"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID2",
			}, http.StatusBadRequest)
			return
		}

		ida, OK := data["ida"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID3",
			}, http.StatusBadRequest)
			return
		}

		//											-->	change here <--
		_, err := query.InsertQuery("tas.option", map[string]interface{}{
			// Data like that
			"idlist_option":   idlo,
			"idlist_activite": idla,
			"idActivite":      ida,
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
