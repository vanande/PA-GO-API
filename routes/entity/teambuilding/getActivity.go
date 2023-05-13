package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
	"time"
)

type Res struct {
	PrixTotal         float64    `json:"prix_total"`
	DateDebut         *time.Time `json:"date_debut"`
	DateFin           *time.Time `json:"date_fin"`
	HeureDebut        *time.Time `json:"heure_debut"`
	HeureFin          *time.Time `json:"heure_fin"`
	IdActivite        int        `json:"idActivite"`
	ActivitePrix      float64    `json:"a_prix"`
	ListActiviteNom   string     `json:"la_nom"`
	ListActiviteDesc  string     `json:"la_description"`
	ListActiviteImage string     `json:"la_image"`
}

func GetActivity(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		// tables in the right order
		tables := []string{"teambuilding_activite  ta ", "activite a", "list_activite la"}
		// columns to look for
		columns := []string{"ta.prix_total", "ta.date_debut", "ta.date_fin", "ta.heure_debut", "ta.heure_fin", "ta.idActivite", "a.prix", "la.nom", "la.description", "la.image"}
		// on what to join
		joins := []map[string]string{
			{"ta.idActivite": "a.idActivite"}, {"a.idList_activite": "la.idList_activite"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"idTEAM_BUILDING": id,
		}

		rows, err := query.SelectWithInnerJoin(tables, columns, joins, conditions)
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		var res []Res
		var prix float64
		prix = 0

		for rows.Next() {
			var s Res
			var dateDebut, dateFin, heureDebut, heureFin []uint8 // need that to deal with the []uint8 values returned by the query
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.PrixTotal, &dateDebut, &dateFin, &heureDebut, &heureFin, &s.IdActivite, &s.ActivitePrix, &s.ListActiviteNom, &s.ListActiviteDesc, &s.ListActiviteImage)
			if err != nil {
				fmt.Println(err)
			}

			// parse []uint8 values into time.Time variables
			if dateDebut != nil {
				t, err := time.Parse("2006-01-02", string(dateDebut))
				if err != nil {
					fmt.Println(err)
				}
				s.DateDebut = &t
			}
			if dateFin != nil {
				t, err := time.Parse("2006-01-02", string(dateFin))
				if err != nil {
					fmt.Println(err)
				}
				s.DateFin = &t
			}
			if heureDebut != nil {
				t, err := time.Parse("15:04:05", string(heureDebut))
				if err != nil {
					fmt.Println(err)
				}
				s.HeureDebut = &t
			}
			if heureFin != nil {
				t, err := time.Parse("15:04:05", string(heureFin))
				if err != nil {
					fmt.Println(err)
				}
				s.HeureFin = &t
			}

			res = append(res, s)
			prix += s.PrixTotal
		}

		libraries.Response(w, map[string]interface{}{
			"message":    "Successfully fetched data",
			"data":       res,
			"prix_total": prix,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
