package place

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

// UpdatePlace updates an existing place in the lieu table in the database, along with its associated address in the adresse table.
func UpdatePlace(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"]
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid place ID",
			}, http.StatusBadRequest)
			return
		}

		fieldToUpdate := make(map[string]interface{})
		if name, OK := data["nom"].(string); OK {
			fieldToUpdate["nom"] = name
		}

		if len(fieldToUpdate) > 0 {

			conditions := map[string]interface{}{
				"idLIEU": id,
			}

			err := query.UpdateQuery("lieu", fieldToUpdate, conditions)
			if err != nil {
				fmt.Println(err)

				libraries.Response(w, map[string]interface{}{
					"message": "Failed to update place",
				}, http.StatusInternalServerError)
				return
			}
		}

		addressFieldToUpdate := make(map[string]interface{})
		if address, OK := data["adresse"].(string); OK {
			addressFieldToUpdate["adresse"] = address
		}
		if zipCode, OK := data["code_postal"].(string); OK {
			addressFieldToUpdate["code_postal"] = zipCode
		}
		if city, OK := data["ville"].(string); OK {
			addressFieldToUpdate["ville"] = city
		}
		if country, OK := data["pays"].(string); OK {
			addressFieldToUpdate["pays"] = country
		}

		if len(addressFieldToUpdate) == 0 && len(fieldToUpdate) == 0 {
			libraries.Response(w, map[string]interface{}{
				"message": "No fields to update",
			}, http.StatusBadRequest)
			return
		}

		addressConditions := map[string]interface{}{
			"idADRESSE": data["ida"],
		}

		err := query.UpdateQuery("adresse", addressFieldToUpdate, addressConditions)
		if err != nil {
			fmt.Println(err)

			libraries.Response(w, map[string]interface{}{
				"message": "Failed to update address",
			}, http.StatusInternalServerError)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Place and address successfully updated",
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
