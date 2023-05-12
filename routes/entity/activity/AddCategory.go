package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func AddCategory(w http.ResponseWriter, req *http.Request) {
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

		//											-->	change here <--
		_, err := query.InsertQuery("category_activite", map[string]interface{}{
			"idCategory":      idc,
			"idlist_activite": ida,
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
