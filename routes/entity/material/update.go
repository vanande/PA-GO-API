package material

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

		id, OK := data["id"].(string)
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

		if description, OK := data["description"].(string); OK {
			fieldToUpdate["description"] = description
		}

		if prix, OK := data["prix"].(string); OK {
			fieldToUpdate["prix"] = prix
		}

		if quantite, OK := data["quantite_disponible"].(string); OK {
			fieldToUpdate["quantite_disponible"] = quantite
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"idMATERIEL": id,
			}

			//                           --> change here <--
			err := query.UpdateQuery("materiel", fieldToUpdate, conditions)
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
