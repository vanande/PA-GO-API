package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func AddAnime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		ida, OK := data["ida"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		idp, OK := data["idp"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		//											-->	change here <--
		_, err := query.InsertQuery("anime", map[string]interface{}{
			// Data like that
			"idlist_activite": ida,
			"idPRESTATAIRE":   idp,
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
