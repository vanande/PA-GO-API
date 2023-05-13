package list

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Res struct {
	IdCategory  int    `json:"idCategory"`
	Nom         string `json:"nom"`
	Description string `json:"description"`
}

func ListActivities(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":

		enableCors(&w)

		data := libraries.Body(w, req)

		id, OK := data["id"].(int)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusUnauthorized)
			return
		}

		// tables in the right order
		tables := []string{"category_activite", "category"}
		// columns to look for
		columns := []string{"ca.idCategory", "c.nom", "c.description"}
		// on what to join
		joins := []map[string]string{
			{"c.idCategory": "ca.idCategory"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"idlist_activite": id,
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
			err := rows.Scan(&s.IdCategory, &s.Nom, &s.Description)
			if err != nil {
				fmt.Println(err)
			}
			res = append(res, s)
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    res,
		}, http.StatusOK)

	case "OPTIONS":
		fmt.Println("Preflight handled")
		w.WriteHeader(http.StatusOK)
	}
}
