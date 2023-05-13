package prestataire

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type S struct {
	IdListActivite string `json:"idlist_activite"`
	IdPrestataire  string `json:"idPRESTATAIRE"`
}

<<<<<<< HEAD
// GetAnime return the activities animated by the presta
=======
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
func GetAnime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		idp, OK := data["idp"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		rows, err := query.SelectQuery("anime", []string{"*"}, map[string]interface{}{"idPRESTATAIRE": idp})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var res []S

		for rows.Next() {
			var s S
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
