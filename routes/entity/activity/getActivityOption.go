package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Res struct {
	IdListOption   int     `json:"idListOption"`
	IdListActivite int     `json:"idListActivite"`
	Prix           float64 `json:"prix"`
	Nom            string  `json:"nom"`
	Description    string  `json:"description"`
}

func GetActivityOption(w http.ResponseWriter, req *http.Request) {
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
		tables := []string{"activite a", "tas.option o", "list_option l"}
		// columns to look for
		columns := []string{"l.idlist_option", "l.idlist_activite", "l.prix", "l.nom", "l.description"}
		// on what to join
		joins := []map[string]string{
			{"a.idActivite": "o.idActivite"},
			{"o.idlist_activite": "l.idlist_activite", "o.idlist_option": "l.idlist_option"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"a.idActivite": id,
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
			err := rows.Scan(&s.IdListOption, &s.IdListActivite, &s.Prix, &s.Nom, &s.Description)
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
