package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
	"time"
)

func AddActivity(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		dateDebutStr := data["date_debut"].(string)
		dateDebut, err := time.Parse("2006-01-02", dateDebutStr)
		if err != nil {
			fmt.Errorf("invalid date: %v", err)
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid date",
			}, http.StatusBadRequest,
			)
		}

		dateFinStr := data["date_fin"].(string)
		dateFin, err := time.Parse("2006-01-02", dateFinStr)
		if err != nil {
			fmt.Errorf("invalid date: %v", err)
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid date",
			}, http.StatusBadRequest,
			)
		}

		//											-->	change here <--
		lastInsertID, err := query.InsertQuery("teambuilding_activite", map[string]interface{}{
			"idActivite":  data["id"],
			"prix_total":  data["prix"],
			"date_debut":  dateDebut,
			"date_fin":    dateFin,
			"heure_debut": data["heure_debut"],
			"heure_fin":   data["heure_fin"],
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
