package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func Update(w http.ResponseWriter, req *http.Request) {
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

		fieldToUpdate := make(map[string]interface{})

		if idc, OK := data["idc"].(string); OK {
			fieldToUpdate["idClient"] = idc
		}

		if t, OK := data["type"].(string); OK {
			fieldToUpdate["type"] = t
		}

		if titre, OK := data["titre"].(string); OK {
			fieldToUpdate["titre"] = titre
		}

		if description, OK := data["description"].(string); OK {
			fieldToUpdate["description"] = description
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"id": id,
			}

			//						-->	change here <--
			err := query.UpdateQuery("team_building", fieldToUpdate, conditions)
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