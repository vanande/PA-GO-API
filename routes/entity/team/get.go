package team

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type S struct {
	IdEquipe       string `json:"ide"`
	IdParticipant  string `json:"idp"`
	IdClient       string `json:"idc"`
	IdTeamBuilding string `json:"idt"`
	Nom            string `json:"nom"`
}

// Get this func get all the users from a team bcs it is the most relevant
func Get(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		ide, OK := data["ide"].(string)
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

		//						-->	change here <--  									-->	and here <--
		rows, err := query.SelectQuery("equipe", []string{"*"}, map[string]interface{}{
			"idEQUIPE":        ide,
			"idCLIENT":        idc,
			"idTEAM_BUILDING": idt,
		})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var res []S

		for rows.Next() {
			var s S
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.IdEquipe, &s.IdParticipant, &s.IdClient, &s.IdTeamBuilding, &s.Nom)
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
