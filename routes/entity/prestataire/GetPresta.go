package prestataire

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Presta struct {
	IdPrestataire  int    `json:"idPRESTATAIRE"`
	Nom            string `json:"nom"`
	Prenom         string `json:"prenom"`
	Tel            string `json:"tel"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Metier         string `json:"metier"`
	Description    string `json:"description"`
	Nom_entreprise string `json:"nom_entreprise"`
	Rib            string `json:"rib"`
	Valide         string `json:"valide"`
	IdAdresse      int    `json:"idADRESSE"`
	Image          string `json:"image"`
}

func GetPresta(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		if data["id"] == nil || data["id"] == "" {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad parameters",
			}, http.StatusBadRequest)
			return
		}

		idPrestataire, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		rows, err := query.SelectQuery("prestataire", []string{"*"}, map[string]interface{}{"idPRESTATAIRE": idPrestataire})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var presta Presta

		for rows.Next() {
			err := rows.Scan(&presta.IdPrestataire, &presta.Nom, &presta.Prenom, &presta.Tel, &presta.Email, &presta.Password, &presta.Metier, &presta.Description, &presta.Nom_entreprise, &presta.Rib, &presta.Valide, &presta.IdAdresse, &presta.Image)
			if err != nil {
				fmt.Println(err)
			}
			presta.Image = fmt.Sprintf("https://%s/public/activity/%s", req.Host, presta.Image)
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    presta,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
