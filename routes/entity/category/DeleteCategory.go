package category

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

<<<<<<< HEAD
		id, OK := data["id"].(string)

		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad ID",
			}, http.StatusNotFound)
			return
		}

		err := query.DeleteQuery("category", "idCategory = ?", id)
=======
		err := query.DeleteQuery("category", "idCategory = ?", data["id"])
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
