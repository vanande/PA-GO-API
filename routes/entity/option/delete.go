package option

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		ida, OK := data["ida"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		//						-->	change here <--  			-->	and here <--
		err := query.DeleteQuery("list_option", "idlist_option = ? AND idlist_activite = ?", id, ida)
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
