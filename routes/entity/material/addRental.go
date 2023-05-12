package material

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
	"time"
)

func AddRental(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		duree, OK := data["duree"].(int)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid duration",
			}, http.StatusBadRequest)
			return
		}
		//											-->	change here <--
		lastInsertID, err := query.InsertQuery("loue", map[string]interface{}{
			// Data like that
			"idt":           data["idt"],
			"idm":           data["idm"],
			"prix_a_payer":  data["prix_a_payer"],
			"date_location": time.Now().Format("2006-01-02"),
			"date_rendu":    time.Now().AddDate(0, 0, duree).Format("2006-01-02"),
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
