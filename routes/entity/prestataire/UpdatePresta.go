package prestataire

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

// UpdatePresta UpdatePlace updates an existing place in the lieu table in the database, along with its associated address in the adresse table.
func UpdatePresta(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"]
		if !OK {
			libraries.Response(w, map[string]interface{}{
<<<<<<< HEAD
				"message": "Invalid place ID",
=======
				"message": "Invalid ID",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			}, http.StatusBadRequest)
			return
		}

		fieldToUpdate := make(map[string]interface{})
		if name, OK := data["nom"].(string); OK {
			fieldToUpdate["nom"] = name
		}

		if prenom, OK := data["prenom"].(string); OK {
			fieldToUpdate["prenom"] = prenom
		}

		if tel, OK := data["tel"].(string); OK {
			fieldToUpdate["tel"] = tel
		}

		if email, OK := data["email"].(string); OK {
			fieldToUpdate["email"] = email
		}

		if password, OK := data["password"].(string); OK {
			fieldToUpdate["password"] = password
		}

		if metier, OK := data["metier"].(string); OK {
			fieldToUpdate["metier"] = metier
		}

		if description, OK := data["description"].(string); OK {
			fieldToUpdate["description"] = description
		}

		if nom_entreprise, OK := data["nom_entreprise"].(string); OK {
			fieldToUpdate["nom_entreprise"] = nom_entreprise
		}

		if rib, OK := data["rib"].(string); OK {
			fieldToUpdate["rib"] = rib
		}

		if valide, OK := data["valide"].(string); OK {
			fieldToUpdate["valide"] = valide
		}

		if idADRESSE, OK := data["idADRESSE"].(string); OK {
			fieldToUpdate["idADRESSE"] = idADRESSE
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"idPRESTATAIRE": id,
			}

			err := query.UpdateQuery("prestataire", fieldToUpdate, conditions)
			if err != nil {
				fmt.Println(err)
				libraries.Response(w, map[string]interface{}{
<<<<<<< HEAD
					"message": "Failed to update place",
=======
					"message": "Failed to update",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
				}, http.StatusInternalServerError)
			}
		}

<<<<<<< HEAD
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
			"message": "Place and address successfully updated",
=======
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
				"message": "Update failed",
			}, http.StatusInternalServerError)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "successfully updated",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
