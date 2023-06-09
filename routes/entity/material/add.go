package material

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func Add(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		//											-->	change here <--
		lastInsertID, err := query.InsertQuery("materiel", map[string]interface{}{
			// Data like that
			"nom":                 data["nom"],
			"description":         data["description"],
			"prix":                data["prix"],
			"quantite_disponible": data["quantite_disponible"],
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
