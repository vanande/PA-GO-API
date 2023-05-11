package company

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Company struct {
	IdClient           string `json:"id"`
	Nom                string `json:"nom"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	Token              string `json:"token"`
	Tel                string `json:"tel"`
	Raison             string `json:"raison_sociale"`
	Creation_du_compte string `json:"creation_du_compte"`
	Pts_fidelite       string `json:"pts_fidelite"`
	IdAdresse          string `json:"idAdresse"`
	Image              string `json:"image"`
}

func GetCompany(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		if data["id"] == nil || data["id"] == "" {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad parameters",
			}, http.StatusBadRequest)
			return
		}

		idCompany, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		rows, err := query.SelectQuery("client", []string{"*"}, map[string]interface{}{"idCLIENT": idCompany})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var company Company

		for rows.Next() {
			err := rows.Scan(&company.IdClient, &company.Nom, &company.Email, &company.Password, &company.Token, &company.Tel, &company.Raison, &company.Creation_du_compte, &company.Pts_fidelite, &company.IdAdresse, &company.Image)
			if err != nil {
				fmt.Println(err)
			}
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    company,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
