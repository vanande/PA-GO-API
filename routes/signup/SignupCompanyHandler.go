package signup

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"net/http"
	"time"
)

func SignupCompany(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)
		if data == nil {
			return
		}

		if data["nom"] == nil ||
			data["email"] == nil ||
			data["telephone"] == nil ||
			data["password"] == nil ||
			data["adresse"] == nil ||
			data["raison_sociale"] == nil ||
			data["statut"] == nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Missing data",
			}, http.StatusBadRequest)
			return
		}

		if !libraries.Correct(data["nom"].(string), "nom") ||
			!libraries.Correct(data["email"].(string), "email") ||
			!libraries.Correct(data["telephone"].(string), "telephone") {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusBadRequest)
			return
		}

		lastInsertId, err := query.InsertQuery("client", map[string]interface{}{
			"nom":                data["nom"],
			"email":              data["email"],
			"tel":                data["telephone"],
			"password":           data["password"],
			"adresse":            data["adresse"],
			"raison_sociale":     data["raison_sociale"],
			"statut":             data["statut"],
			"creation_du_compte": time.Now(),
			"pts_fidelite":       "0",
		})
		if err != nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Insert failed",
			}, http.StatusBadRequest)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Insert success",
			"id":      lastInsertId,
		}, http.StatusOK)
	}
}
