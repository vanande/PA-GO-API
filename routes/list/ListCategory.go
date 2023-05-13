package list

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

func ListCategory(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		rows, err := query.SelectQuery("category", []string{"*"}, map[string]interface{}{"1": 1})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var res []S

		for rows.Next() {
			var s S
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
