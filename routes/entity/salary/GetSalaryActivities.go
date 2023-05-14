package salary

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type ListActivite struct {
	ID            string
	Nom           string
	Description   string
	NbPersonneMin string
	NbPersonneMax string
	PrixMin       string
	PrixMax       string
	Image         string
	Note          string
	HeureDebut    string
	HeureFin      string
	DateDebut     string
	DateFin       string
}

func GetSalaryActivities(w http.ResponseWriter, req *http.Request) {
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

		idc, OK := data["idc"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		// tables in the right order
		tables := []string{"equipe e", "team_building tb", "teambuilding_activite tba", "activite a", "list_activite la"}
		// columns to look for
		columns := []string{"la.*", "tba.heure_debut", "tba.heure_fin", "tba.date_debut", "tba.date_fin"}
		// on what to join
		joins := []map[string]string{
			{"e.idTEAM_BUILDING": "tb.idTEAM_BUILDING"},
			{"tb.idTEAM_BUILDING": "tba.idTEAM_BUILDING"},
			{"tba.idACTIVITE": "a.idACTIVITE"},
			{"a.idlist_activite": "la.idlist_activite"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"e.idPARTICIPANT": id,
			"e.idCLIENT":      idc,
		}

		rows, err := query.SelectWithInnerJoin(tables, columns, joins, conditions)
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		var res []ListActivite

		for rows.Next() {
			var a ListActivite
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&a.ID, &a.Nom, &a.Description, &a.NbPersonneMin, &a.NbPersonneMax, &a.PrixMin, &a.PrixMax, &a.Image, &a.Note, &a.HeureDebut, &a.HeureFin, &a.DateDebut, &a.DateFin)
			a.Image = fmt.Sprintf("https://togetherandstronger.fr:9000/public/activity/%s", a.Image)
			if err != nil {
				fmt.Println(err)
			}
			res = append(res, a)
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    res,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
