package prestataire

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

// AddPresta AddPlace create a new place and the associated address in the lieu and adresse tables in the database.
func AddPresta(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		lastAddressID, err := query.InsertQuery("adresse", map[string]interface{}{
			"adresse":     data["adresse"],
			"complement":  data["complement"],
			"ville":       data["ville"],
			"code_postal": data["code_postal"],
			"pays":        data["pays"],
		})

		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		lastInsertID, err := query.InsertQuery("prestataire", map[string]interface{}{
			"nom":            data["nom"],
			"prenom":         data["prenom"],
			"tel":            data["tel"],
			"email":          data["email"],
			"password":       data["password"],
			"metier":         data["metier"],
			"description":    data["description"],
			"nom_entreprise": data["nom_entreprise"],
			"idADRESSE":      lastAddressID,
			"rib":            "",
			"valide":         "1",
		})
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "Bad request",
			}, http.StatusBadRequest)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message":             "Successfully added",
			"lastAddressInserted": lastAddressID,
			"lastPrestaInserted":  lastInsertID,
		}, http.StatusOK)
	}
}
