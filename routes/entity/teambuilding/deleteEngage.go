package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func DeleteEngage(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

<<<<<<< HEAD
		id, OK := data["idt"].(string)
=======
		id, OK := data["id"].(string)
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
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
		err := query.DeleteQuery("engage", "idTEAM_BUILDING = ? AND idPRESTATAIRE = ?", id, idp)
=======
		err := query.DeleteQuery("engage", "idTEAM_BUILDING = ? AND idPRESTATAIRE", id, idp)
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
