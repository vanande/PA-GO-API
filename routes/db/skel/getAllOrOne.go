package skel

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"log"
	"net/http"
)

type Adresse struct {
	IdAdresse   string `json:"idADRESSE"`
	Adresse     string `json:"adresse"`
	Complement  string `json:"complement"`
	Ville       string `json:"ville"`
	Code_postal string `json:"code_postal"`
	Pays        string `json:"pays"`
}

type Presta struct {
	IdPrestataire  string  `json:"idPRESTATAIRE"`
	Nom            string  `json:"nom"`
	Prenom         string  `json:"prenom"`
	Tel            string  `json:"tel"`
	Email          string  `json:"email"`
	Metier         string  `json:"metier"`
	Description    string  `json:"description"`
	Nom_entreprise string  `json:"nom_entreprise"`
	Rib            string  `json:"rib"`
	Valide         string  `json:"valide"`
	AdressePresta  Adresse `json:"adressePresta"`
}

func GetAllOrOne(w http.ResponseWriter, req *http.Request) {
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

		// tables in the right order
		tables := []string{"prestataire p", "adresse a"}
		// columns to look for
		columns := []string{"p.idPRESTATAIRE", "p.nom", "p.prenom", "p.tel", "p.email", "p.metier", "p.description", "p.nom_entreprise", "p.rib", "p.valide", "a.adresse", "a.complement", "a.ville", "a.code_postal", "a.pays", "a.idADRESSE"}
		// on what to join
		joins := []map[string]string{
			{"p.idADRESSE": "a.idADRESSE"},
		}
		// the where at the end

		var conditions map[string]interface{}

		if id == "*" {
			conditions = map[string]interface{}{}
		} else {
			conditions = map[string]interface{}{
				"idPRESTATAIRE": id,
			}
		}

		rows, err := query.SelectWithInnerJoin(tables, columns, joins, conditions)
		if err != nil {
			log.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		var res []Presta

		for rows.Next() {
			var s Presta
			var a Adresse
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.IdPrestataire, &s.Nom, &s.Prenom, &s.Tel, &s.Email, &s.Metier, &s.Description, &s.Nom_entreprise, &s.Rib, &s.Valide, &a.Adresse, &a.Complement, &a.Ville, &a.Code_postal, &a.Pays, &a.IdAdresse)
			if err != nil {
				log.Println(err)
			}

			s.AdressePresta = a
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
