package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
	"time"
)

type Materiel struct {
	IdMateriel         string     `json:"id"`
	Nom                string     `json:"nom"`
	Description        string     `json:"description"`
	Prix               float64    `json:"prix"`
	QuantiteDisponible int        `json:"quantite_disponible"`
	PrixAPayer         float64    `json:"prix_a_payer"`
	DateLocation       *time.Time `json:"date_location"`
	DateRendu          *time.Time `json:"date_rendu"`
}

func GetMateriel(w http.ResponseWriter, req *http.Request) {
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
		tables := []string{"loue l", "materiel m"}
		// columns to look for
		columns := []string{"m.idMATERIEL", "m.nom", "m.description", "m.prix", "m.quantite_disponible", "l.prix_a_payer", "l.date_location", "l.date_rendu"}
		// on what to join
		joins := []map[string]string{
			{"l.idMATERIEL": "m.idMATERIEL"},
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

		var res []Materiel
		var prix float64

		for rows.Next() {
			var m Materiel
			var DateLocation, DateRendu []uint8
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&m.IdMateriel, &m.Nom, &m.Description, &m.Prix, &m.QuantiteDisponible, &m.PrixAPayer, &DateLocation, &DateRendu)
			if err != nil {
				fmt.Println(err)
			}

			if DateLocation != nil {
				t, _ := time.Parse("2006-01-02", string(DateLocation))
				m.DateLocation = &t
			}

			if DateRendu != nil {
				t, _ := time.Parse("2006-01-02", string(DateRendu))
				m.DateRendu = &t
			}

			prix += m.PrixAPayer
			res = append(res, m)
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
