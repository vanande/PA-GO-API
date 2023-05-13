package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func AddEngage(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		//											-->	change here <--
		_, err := query.InsertQuery("engage", map[string]interface{}{
			// Data like that
			"idTEAM_BUILDING": data["idt"],
			"idPrestataire":   data["idp"],
			"prix_a_payer":    data["prix"],
		})
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully added",
		}, http.StatusOK)
	}
}
