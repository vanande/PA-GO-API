package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Salle struct {
	IdSalle      string  `json:"id"`
	PrixAPayer   float64 `json:"PrixAPayer"`
	Nom          string  `json:"nom"`
	NumSalles    string  `json:"numSalles"`
	Prix         float64 `json:"prix"`
	Description  string  `json:"description"`
	Disponibilte string  `json:"disponibilite"`
	IdLieu       string  `json:"idLieu"`
}

func GetSalle(w http.ResponseWriter, req *http.Request) {
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
		tables := []string{"reserve r", "salle s"}
		// columns to look for
		columns := []string{"r.idSalle", "r.prix_a_payer", "s.nom", "s.num_salle", "s.prix", "s.description", "s.disponibilite", "s.idLieu"}
		// on what to join
		joins := []map[string]string{
			{"r.idSALLE": "s.idSALLE"},
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

		var res []Salle
		var prix float64
		for rows.Next() {
			var s Salle
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.IdSalle, &s.PrixAPayer, &s.Nom, &s.NumSalles, &s.Prix, &s.Description, &s.Disponibilte, &s.IdLieu)
			if err != nil {
				fmt.Println(err)
			}
			res = append(res, s)
			prix += s.PrixAPayer
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    res,
			"prix":    prix,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
