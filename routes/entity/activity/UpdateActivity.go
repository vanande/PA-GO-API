package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

// UpdateActivity updates an existing activity in the list_activite table in the database.
//
// w: the http.ResponseWriter used to send the response
// req: the *http.Request object representing the incoming request
//
// Returns nothing, but sends an HTTP response indicating whether the activity was successfully updated.
func UpdateActivity(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"]
		if !OK {
			libraries.Response(w, map[string]interface{}{
<<<<<<< HEAD
				"message": "Invalid place ID",
=======
				"message": "Invalid activity ID",
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			}, http.StatusBadRequest)
			return
		}

		fieldToUpdate := make(map[string]interface{})
		if name, OK := data["name"].(string); OK {
			fieldToUpdate["name"] = name
		}
		if description, err := data["description"].(string); err {
			fieldToUpdate["description"] = description
		}
<<<<<<< HEAD
		if minParticipants, err := data["min_participants"].(string); err {
			fieldToUpdate["nb_personne_min"] = minParticipants
		}
		if maxParticipants, err := data["max_participants"].(string); err {
			fieldToUpdate["nb_personne_max"] = maxParticipants
		}
		if minPrice, err := data["min_price"].(string); err {
			fieldToUpdate["min_price"] = minPrice
		}
		if maxPrice, err := data["max_price"].(string); err {
=======
		if minParticipants, err := data["min_participants"].(float64); err {
			fieldToUpdate["min_participants"] = minParticipants
		}
		if maxParticipants, err := data["max_participants"].(float64); err {
			fieldToUpdate["max_participants"] = maxParticipants
		}
		if minPrice, err := data["min_price"].(float64); err {
			fieldToUpdate["min_price"] = minPrice
		}
		if maxPrice, err := data["max_price"].(float64); err {
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			fieldToUpdate["max_price"] = maxPrice
		}
		if image, err := data["image"].(string); err {
			fieldToUpdate["image"] = image
		}

		if len(fieldToUpdate) == 0 {
			libraries.Response(w, map[string]interface{}{
				"message": "No valid fields to update",
			}, http.StatusBadRequest)
			return
		}

		conditions := map[string]interface{}{
			"idlist_activite": id,
		}

		err := query.UpdateQuery("list_activite", fieldToUpdate, conditions)
		if err != nil {
			fmt.Println(err)

			libraries.Response(w, map[string]interface{}{
				"message": "Failed to update activity",
			}, http.StatusInternalServerError)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Activity successfully updated",
		}, http.StatusOK)
	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
