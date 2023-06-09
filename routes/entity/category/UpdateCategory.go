package category

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

<<<<<<< HEAD
func Update(w http.ResponseWriter, req *http.Request) {
=======
func UpdateCategory(w http.ResponseWriter, req *http.Request) {
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
	switch req.Method {
	case "POST":
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

		if description, OK := data["description"].(string); OK {
			fieldToUpdate["description"] = description
		}

		if len(fieldToUpdate) > 0 {
			conditions := map[string]interface{}{
				"idCategory": id,
			}

			err := query.UpdateQuery("category", fieldToUpdate, conditions)
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
