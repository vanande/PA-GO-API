package category

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

<<<<<<< HEAD
func Add(w http.ResponseWriter, req *http.Request) {
=======
func AddCategory(w http.ResponseWriter, req *http.Request) {
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

<<<<<<< HEAD
		lastInsertID, err := query.InsertQuery("category", map[string]interface{}{
=======
		lastInsertID, err := query.InsertQuery("client", map[string]interface{}{
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			"nom":         data["nom"],
			"description": data["description"],
		})
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message":        "Successfully added",
			"lastIDInserted": lastInsertID,
		}, http.StatusOK)
	}
}
