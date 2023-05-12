package company

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func UpdateCompany(w http.ResponseWriter, req *http.Request) {
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

		if email, OK := data["email"].(string); OK {
			fieldToUpdate["email"] = email
		}

		if password, OK := data["password"].(string); OK {
			fieldToUpdate["password"] = password
		}

		if token, OK := data["token"].(string); OK {
			fieldToUpdate["token"] = token
		}

		if tel, OK := data["tel"].(string); OK {
			fieldToUpdate["tel"] = tel
		}

		if raison_sociale, OK := data["raison_sociale"].(string); OK {
			fieldToUpdate["raison_sociale"] = raison_sociale
		}

		if pts_fidelite, OK := data["pts_fidelite"].(string); OK {
			fieldToUpdate["pts_fidelite"] = pts_fidelite
		}

		if idADRESSE, OK := data["idADRESSE"].(string); OK {
			fieldToUpdate["idADRESSE"] = idADRESSE
		}

		if image, OK := data["image"].(string); OK {
			fieldToUpdate["image"] = image
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"idCLIENT": id,
			}

			err := query.UpdateQuery("client", fieldToUpdate, conditions)
			if err != nil {
				fmt.Println(err)
				libraries.Response(w, map[string]interface{}{
					"message": "Failed to update place",
				}, http.StatusInternalServerError)
			}
		}

		ida, OK := data["ida"]
		if OK {
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
				"idADRESSE": ida,
			}

			err := query.UpdateQuery("adresse", addressFieldToUpdate, addressConditions)
			if err != nil {
				fmt.Println(err)
				libraries.Response(w, map[string]interface{}{
					"message": "Failed to update address",
				}, http.StatusInternalServerError)
				return
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
