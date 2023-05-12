package prestataire

import (
	"TogetherAndStronger/libraries"
	db2 "TogetherAndStronger/routes/db/init"
	"fmt"
	"html"
	"log"
	"net/http"
)

type Presta struct {
	Nom            string `json:"nom"`
	Prenom         string `json:"prenom"`
	Tel            string `json:"tel"`
	Email          string `json:"email"`
	Metier         string `json:"metier"`
	Description    string `json:"description"`
	Nom_entreprise string `json:"nom_entreprise"`
	Rib            string `json:"rib"`
	Valide         string `json:"valide"`
	Adresse        string `json:"adresse"`
	Complement     string `json:"complement"`
	Ville          string `json:"ville"`
	Code_postal    string `json:"code_postal"`
	Pays           string `json:"pays"`
	IdAdresse      int    `json:"idADRESSE"`
}

func GetPresta(w http.ResponseWriter, req *http.Request) {
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

		idPrestataire, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		db, err := db2.InitDB()
		if err != nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Error connecting to database",
			}, http.StatusInternalServerError)
			log.Fatalf("Error connecting to database : %v", err)
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Fatalf("Error closing database : %v", err)
			}
		}()

		query := fmt.Sprintf(`
			SELECT 
				p.nom, p.prenom, p.tel, p.email, p.metier, p.description, p.nom_entreprise, p.rib, p.valide, a.adresse, a.complement, a.ville, a.code_postal, a.pays
			FROM 
				prestataire p 
				INNER JOIN adresse a ON a.idADRESSE = p.idADRESSE
			WHERE p.idPRESTATAIRE = %s`,
			idPrestataire)
		log.Printf("SQL query: %s", html.EscapeString(query))

		presta := Presta{}

		rows, err := db.Query(query)
		if err != nil {
			log.Printf("Error querying database: %v", err)
			libraries.Response(w, map[string]interface{}{
				"message": "Error querying database",
			}, http.StatusInternalServerError)
			return
		}
		defer func() {
			if err := rows.Close(); err != nil {
				log.Printf("Error closing rows: %v", err)
			}
		}()
		if rows.Next() {
			if err := rows.Scan(&presta.Nom, &presta.Prenom, &presta.Tel, &presta.Email, &presta.Metier, &presta.Description, &presta.Nom_entreprise, &presta.Rib, &presta.Valide, &presta.Adresse, &presta.Complement, &presta.Ville, &presta.Code_postal, &presta.Pays); err != nil {
				log.Printf("Error scanning row: %v", err)
				libraries.Response(w, map[string]interface{}{
					"message": "Error querying database",
				}, http.StatusInternalServerError)
				return
			}
			presta.IdAdresse = 0
			libraries.Response(w, map[string]interface{}{
				"prestataire": presta,
			}, http.StatusOK)
		} else {
			libraries.Response(w, map[string]interface{}{
				"message": "Prestataire not found",
			}, http.StatusNotFound)
		}
	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
