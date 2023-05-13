package activity

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type S struct {
	IdCategory  string `json:"id"`
	Nom         string `json:"nom"`
	Description string `json:"description"`
}

func GetCategory(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		if data["id"] == nil || data["id"] == "" {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad parameters",
			}, http.StatusBadRequest)
			return
		}

		id, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

<<<<<<< HEAD
		tables := []string{"category_activite ca", "category c"}
=======
		tables := []string{"category_activite", "category"}
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
		columns := []string{"ca.idCategory", "c.nom", "c.description"}
		joins := []map[string]string{
			{"c.idCategory": "ca.idCategory"},
		}
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

		var res []S

		for rows.Next() {
			var s S
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

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
