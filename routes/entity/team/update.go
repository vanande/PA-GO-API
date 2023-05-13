package team

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func Update(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PATCH":
		data := libraries.Body(w, req)

		ide, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		idc, OK := data["idc"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}
		idt, OK := data["idt"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		fieldToUpdate := make(map[string]interface{})

		// as much field as there is updatable cols in the table here

		if name, OK := data["nom"].(string); OK {
			fieldToUpdate["nom"] = name
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"idEQUIPE":        ide,
				"idCLIENT":        idc,
				"idTEAM_BUILDING": idt,
			}

			//                           --> change here <--
			err := query.UpdateQuery("equipe", fieldToUpdate, conditions)
			if err != nil {
				fmt.Println(err)
				libraries.Response(w, map[string]interface{}{
					"message": "Failed to update place",
				}, http.StatusInternalServerError)
			}
		}

		libraries.Response(w, map[string]interface{}{
			"message": "successfully updated",
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
