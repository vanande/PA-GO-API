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
<<<<<<< HEAD
				"idTEAM_BUILDING": id,
			}

			//			-->	change here <--
=======
				"id": id,
			}

			//						-->	change here <--
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			err := query.UpdateQuery("team_building", fieldToUpdate, conditions)
			if err != nil {
				fmt.Println(err)
				libraries.Response(w, map[string]interface{}{
<<<<<<< HEAD
					"message": "Failed to update",
				}, http.StatusInternalServerError)
				return
=======
					"message": "Failed to update place",
				}, http.StatusInternalServerError)
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
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
