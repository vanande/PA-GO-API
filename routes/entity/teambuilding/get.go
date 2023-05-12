package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type S struct {
	IdTeamBuilding string `json:"id"`
	IdClient       string `json:"idc"`
	Type           string `json:"type"`
	Titre          string `json:"titre"`
	Description    string `json:"description"`
}

func Get(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad parameters",
			}, http.StatusBadRequest)
			return
		}

		rows, err := query.SelectQuery("s", []string{"*"}, map[string]interface{}{"id": id})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var s S

		for rows.Next() {
			err := rows.Scan(&s.IdTeamBuilding, &s.IdClient, &s.Type, &s.Titre, &s.Description)
			if err != nil {
				fmt.Println(err)
			}
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    s,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
