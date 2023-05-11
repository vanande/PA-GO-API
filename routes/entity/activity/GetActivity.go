package activity

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
}

func GetActivity(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		if data["id"] == nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Empty parameters",
			}, http.StatusBadRequest)
			return
		}
		if data["id"] == "" {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		rows, err := query.SelectQuery("list_activite", []string{"*"}, map[string]interface{}{"idlist_activite": data["id"]})
		if err != nil {
			fmt.Errorf("Select failed : %v", err)
		}

		var listActivites []ListActivite

		for rows.Next() {
			var la ListActivite
			err := rows.Scan(&la.ID, &la.Nom, &la.Description, &la.NbPersonneMin, &la.NbPersonneMax, &la.PrixMin, &la.PrixMax, &la.Image)
			if err != nil {
				fmt.Println(err)
			}
			la.Image = fmt.Sprintf("https://%s/public/activity/%s", req.Host, la.Image)
			listActivites = append(listActivites, la)
		}
		fmt.Println(listActivites)

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    listActivites,
		}, http.StatusOK)
	}
}
