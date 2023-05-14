package salary

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"log"
	"net/http"
)

type Salary struct {
	Id        string `json:"id"`
	Nom       string `json:"nom"`
	Prenom    string `json:"prenom"`
	Tel       string `json:"tel"`
	Age       string `json:"age"`
	Sexe      string `json:"sexe"`
	IdCompany string `json:"idCompany"`
}

func GetSalary(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		libraries.EnableCors(&w)

		data := libraries.Body(w, req)

		id, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		condition := map[string]interface{}{}
		if id == "*" {
			condition = nil
		} else {
			idc, OK := data["idc"].(string)
			if !OK {
				libraries.Response(w, map[string]interface{}{
					"message": "Invalid parameters",
				}, http.StatusBadRequest)
				return
			}
			condition["idPARTICIPANT"] = id
			condition["idCLIENT"] = idc
		}

		rows, err := query.SelectQuery("participant", []string{"idPARTICIPANT", "nom", "prenom", "telephone", "age", "sexe", "idCLIENT"}, condition)
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var res []Salary

		for rows.Next() {
			var s Salary
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.Id, &s.Nom, &s.Prenom, &s.Tel, &s.Age, &s.Sexe, &s.IdCompany)
			if err != nil {
				log.Println(err)
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
