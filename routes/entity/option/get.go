package option

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type S struct {
	IdListOption   string `json:"id"`
	IdListActivite string `json:"ida"`
	Nom            string `json:"nom"`
	Prix           string `json:"prix"`
	Description    string `json:"description"`
}

func Get(w http.ResponseWriter, req *http.Request) {
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

		ida, OK := data["ida"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		//						-->	change here <--  									-->	and here <--
		rows, err := query.SelectQuery("list_option", []string{"*"}, map[string]interface{}{
			"idlist_option":   id,
			"idlist_activite": ida,
		})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var s S

		for rows.Next() {
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.IdListOption, &s.IdListActivite, &s.Nom, &s.Prix, &s.Description)
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
