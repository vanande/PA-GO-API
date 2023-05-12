package room

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func DeleteRoom(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		err := query.DeleteQuery("salle", "idSALLE = ?", data["id"])
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
