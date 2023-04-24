package signup

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"net/http"
)

func SignupPresta(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)
		if data == nil {
			return
		}

		if data["nom"] == nil ||
			data["prenom"] == nil ||
			data["email"] == nil ||
			data["telephone"] == nil ||
			data["password"] == nil ||
			data["metier"] == nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Missing data",
			}, http.StatusBadRequest)
			return
		}

		if !libraries.Correct(data["nom"].(string), "nom") ||
			!libraries.Correct(data["prenom"].(string), "nom") ||
			!libraries.Correct(data["email"].(string), "email") ||
			!libraries.Correct(data["telephone"].(string), "telephone") {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusBadRequest)
			return
		}

		lastInsertId, err := query.InsertQuery("prestataire", map[string]interface{}{
			"nom":         data["nom"],
			"prenom":      data["prenom"],
			"email":       data["email"],
			"tel":         data["telephone"],
			"password":    data["password"],
			"metier":      data["metier"],
			"description": data["description"],
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
	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
