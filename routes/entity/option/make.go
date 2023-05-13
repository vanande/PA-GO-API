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
<<<<<<< HEAD
				"message": "Invalid ID1",
=======
				"message": "Invalid ID",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			}, http.StatusBadRequest)
			return
		}

		idla, OK := data["idla"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
<<<<<<< HEAD
				"message": "Invalid ID2",
=======
				"message": "Invalid ID",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			}, http.StatusBadRequest)
			return
		}

		ida, OK := data["ida"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
<<<<<<< HEAD
				"message": "Invalid ID3",
=======
				"message": "Invalid ID",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			}, http.StatusBadRequest)
			return
		}

		//											-->	change here <--
<<<<<<< HEAD
		_, err := query.InsertQuery("tas.option", map[string]interface{}{
			// Data like that
=======
		lastInsertID, err := query.InsertQuery("tas.option", map[string]interface{}{
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
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
<<<<<<< HEAD
			"message": "Successfully added",
=======
			"message":        "Successfully added",
			"lastIDInserted": lastInsertID,
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
		}, http.StatusOK)
	}
}
