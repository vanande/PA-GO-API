package room

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func AddRoom(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		lastInsertID, err := query.InsertQuery("salle", map[string]interface{}{
			"nom":           data["nom"],
			"num_salle":     data["num_salle"],
			"prix":          data["prix"],
			"description":   data["description"],
			"disponibilite": 1,
			"idLIEU":        data["idl"],
			"image":         data["image"],
		})
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message":          "Successfully added",
			"lastRoomInserted": lastInsertID,
		}, http.StatusOK)
	}
}
