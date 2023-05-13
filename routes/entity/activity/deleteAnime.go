package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func DeleteAnime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		ida, OK := data["ida"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		idp, OK := data["idp"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		//						-->	change here <--  			-->	and here <--
<<<<<<< HEAD
		err := query.DeleteQuery("anime", "idlist_activite = ? AND idPRESTATAIRE = ?", ida, idp)
=======
		err := query.DeleteQuery("", "idlist_activite = ? AND idPRESTATAIRE = ?", ida, idp)
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
		if err != nil {
			fmt.Println(w, err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully deleted data",
		}, http.StatusOK)
	}
}
