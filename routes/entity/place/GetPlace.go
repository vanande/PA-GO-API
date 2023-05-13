package place

import (
	"TogetherAndStronger/libraries"
	db2 "TogetherAndStronger/routes/db/init"
	"fmt"
	"html"
	"net/http"
)

type Place struct {
	Nom         string `json:"nom"`
	IdAdresse   int    `json:"idADRESSE"`
	Adresse     string `json:"adresse"`
	Complement  string `json:"complement"`
	Ville       string `json:"ville"`
	Code_postal string `json:"code_postal"`
	Pays        string `json:"pays"`
}

func GetPlace(w http.ResponseWriter, req *http.Request) {
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

		idLIEU := data["id"]

		db, err := db2.InitDB()
		if err != nil {
			fmt.Errorf("Error connecting to database : %v", err)
		}
		defer db.Close()

		query := fmt.Sprintf(`
			SELECT 
				l.nom, l.idADRESSE, a.adresse, a.complement, a.ville, a.code_postal, a.pays
			FROM 
				lieu l 
				INNER JOIN adresse a ON a.idADRESSE = l.idADRESSE
<<<<<<< HEAD
			WHERE l.idLIEU = %s`,
=======
			WHERE l.id = %s`,
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			idLIEU)
		fmt.Println(html.EscapeString(query))

		place := Place{}

		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		}

		for rows.Next() {
			err := rows.Scan(&place.Nom, &place.IdAdresse, &place.Adresse, &place.Complement, &place.Ville, &place.Code_postal, &place.Pays)
			if err != nil {
				fmt.Println(err)
			}
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    place,
		}, http.StatusOK)
	}
}
