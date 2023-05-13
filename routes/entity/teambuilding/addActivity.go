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

<<<<<<< HEAD
		dateDebutStr := data["date_debut"].(string)
		dateDebut, err := time.Parse("2006-01-02", dateDebutStr)
=======
		dateDebut, err := time.Parse("2006-01-02", data["date_debut"].(string))
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
		if err != nil {
			fmt.Errorf("invalid date: %v", err)
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid date",
			}, http.StatusBadRequest,
			)
		}

<<<<<<< HEAD
		dateFinStr := data["date_fin"].(string)
		dateFin, err := time.Parse("2006-01-02", dateFinStr)
=======
		dateFin, err := time.Parse("2006-01-02", data["date_fin"].(string))
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
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
