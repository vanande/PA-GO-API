package skel

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Res struct {
	IdParticipant string `json:"id"`
	Nom           string `json:"nom"`
	Prenom        string `json:"prenom"`
	Tel           string `json:"tel"`
	Age           string `json:"age"`
	Sexe          string `json:"sexe"`
	IdClient      string `json:"idc"`
}

func GetSub(w http.ResponseWriter, req *http.Request) {
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

		idt, OK := data["idt"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		// tables in the right order
		tables := []string{"equipe e", "participant p"}
		// columns to look for
		columns := []string{"p.idPARTICIPANT", "p.nom", "p.prenom", "p.tel", "p.age", "p.sexe", "e.idCLIENT"}
		// on what to join
		joins := []map[string]string{
			{"e.idPARTICIPANT": "p.idPARTICIPANT"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"idEQUIPE":        id,
			"idCLIENT":        idc,
			"idTEAM_BUILDING": idt,
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
			err := rows.Scan(&s.IdParticipant, &s.Nom, &s.Prenom, &s.Tel, &s.Age, &s.Sexe, &s.IdClient)
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
