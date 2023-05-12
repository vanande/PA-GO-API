package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Anime struct {
	IdListActivite string `json:"idlist_activite"`
	IdPrestataire  string `json:"idPRESTATAIRE"`
}

func GetAnime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		ida, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		rows, err := query.SelectQuery("anime", []string{"*"}, map[string]interface{}{"idlist_activite": ida})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var res []Anime

		for rows.Next() {
			var s Anime
			err := rows.Scan(&s.IdListActivite, &s.IdPrestataire)
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
