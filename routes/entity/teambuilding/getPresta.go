package teambuilding

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Presta struct {
	IdPrestataire string  `json:"id"`
	Nom           string  `json:"nom"`
	Prenom        string  `json:"prenom"`
	Tel           string  `json:"tel"`
	Email         string  `json:"email"`
	Metier        string  `json:"metier"`
	Description   string  `json:"description"`
	NomEntreprise string  `json:"nomEntreprise"`
	Rib           string  `json:"rib"`
	Valide        string  `json:"valide"`
	IdAdresse     string  `json:"idAdresse"`
	Image         string  `json:"image"`
	Prix          float64 `json:"prix"`
}

func GetPresta(w http.ResponseWriter, req *http.Request) {
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

		// tables in the right order
		tables := []string{"engage e", "prestataire p"}
		// columns to look for
		columns := []string{"p.idPRESTATAIRE", "p.nom", "p.prenom", "p.tel", "p.email", "p.metier", "p.description", "p.nom_entreprise", "p.rib", "p.valide", "p.idADRESSE", "p.image", "e.prix_a_payer"}
		// on what to join
		joins := []map[string]string{
			{"e.idPRESTATAIRE": "p.idPRESTATAIRE"},
		}
		// the where at the end
		conditions := map[string]interface{}{
			"idTEAM_BUILDING": id,
		}

		rows, err := query.SelectWithInnerJoin(tables, columns, joins, conditions)
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		var res []Presta
		var prix float64
<<<<<<< HEAD
		prix = 0
=======
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c

		for rows.Next() {
			var p Presta
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&p.IdPrestataire, &p.Nom, &p.Prenom, &p.Tel, &p.Email, &p.Metier, &p.Description, &p.NomEntreprise, &p.Rib, &p.Valide, &p.IdAdresse, &p.Image, &p.Prix)
			if err != nil {
				fmt.Println(err)
			}
			prix += p.Prix
			res = append(res, p)
		}

		libraries.Response(w, map[string]interface{}{
			"message":    "Successfully fetched data",
			"data":       res,
			"prix_total": prix,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
