package room

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func UpdateRoom(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PATCH":
		data := libraries.Body(w, req)

		id, OK := data["id"]
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid ID",
			}, http.StatusBadRequest)
			return
		}

		fieldToUpdate := make(map[string]interface{})
		if name, OK := data["nom"].(string); OK {
			fieldToUpdate["nom"] = name
		}

		if numSalle, OK := data["num_salle"].(string); OK {
			fieldToUpdate["num_salle"] = numSalle
		}

		if prix, OK := data["prix"].(string); OK {
			fieldToUpdate["prix"] = prix
		}

		if description, OK := data["description"].(string); OK {
			fieldToUpdate["description"] = description
		}

		if disponibilite, OK := data["disponibilite"].(string); OK {
			fieldToUpdate["disponibilite"] = disponibilite
		}

		if idLIEU, OK := data["idLIEU"].(string); OK {
			fieldToUpdate["idLIEU"] = idLIEU
		}

		if image, OK := data["image"].(string); OK {
			fieldToUpdate["image"] = image
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"idSALLE": id,
			}

			err := query.UpdateQuery("salle", fieldToUpdate, conditions)
			if err != nil {
				fmt.Println(err)
				libraries.Response(w, map[string]interface{}{
					"message": "Failed to update",
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
