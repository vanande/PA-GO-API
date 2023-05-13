package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func DeleteEngage(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "DELETE":
		data := libraries.Body(w, req)

		id, OK := data["idt"].(string)
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
		err := query.DeleteQuery("engage", "idTEAM_BUILDING = ? AND idPRESTATAIRE = ?", id, idp)
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
