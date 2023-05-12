package room

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Res struct {
	IdListActivite string `json:"ida"`
	Nom            string `json:"nom"`
	Description    string `json:"description"`
	Personne_min   string `json:"personne_min"`
	Personne_max   string `json:"personne_max"`
	Prix_min       string `json:"prix_min"`
	Prix_max       string `json:"prix_max"`
	Image          string `json:"image"`
}

//	select a.idlist_activite, a.nom, a.description, a.nb_personne_min, a.nb_personne_max, a.prix_min, a.prix_max, a.image
//
// from activite_possible ap
// INNER JOIN list_activite a ON a.idlist_activite = ap.idlist_activite
// WHERE idSALLE = 1;
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
		tables := []string{"activite_possible ap", "list_activite a"}
		// columns to look for
		columns := []string{"a.idlist_activite", "a.nom", "a.description", "a.nb_personne_min", "a.nb_personne_max", "a.prix_min", "a.prix_max", "a.image"}
		// on what to join
		joins := []map[string]string{
			{"a.idlist_activite": "ap.idlist_activite"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"idSALLE": id,
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

		for rows.Next() {
			var s Res
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.IdListActivite, &s.Nom, &s.Description, &s.Personne_min, &s.Personne_max, &s.Prix_min, &s.Prix_max, &s.Image)
			if err != nil {
				fmt.Println(err)
			}
			res = append(res, s)
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
