package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func DeleteCategory(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		idc, OK := data["idc"].(string)
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
		err := query.DeleteQuery("category_activite", "idCategory = ? AND idlist_activite = ?", idc, ida)
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
