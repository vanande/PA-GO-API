package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

// AddActivity adds a new activity to the list_activite table in the database.
//
// w: the http.ResponseWriter used to send the response
// req: the *http.Request object representing the incoming request
//
// Returns nothing, but sends an HTTP response indicating whether the activity was successfully added.
func AddActivity(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		_, err := query.InsertQuery("list_activite", map[string]interface{}{
			"nom":             data["nom"],
			"description":     data["description"],
			"nb_personne_min": data["nb_personne_min"],
			"nb_personne_max": data["nb_personne_max"],
			"prix_min":        data["prix_min"],
			"prix_max":        data["prix_max"],
			"image":           data["image"],
		})

		if err != nil {
			fmt.Println(err)

			libraries.Response(w, map[string]interface{}{
				"message": "Duplicate entry",
			}, http.StatusBadRequest)
			return
		}

	}
	libraries.Response(w, map[string]interface{}{
		"message": "Successfully added",
	}, http.StatusOK)

}
