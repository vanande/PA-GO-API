package company

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
	"time"
)

func AddCompany(w http.ResponseWriter, req *http.Request) {
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

		lastInsertID, err := query.InsertQuery("client", map[string]interface{}{
			"nom":                data["nom"],
			"email":              data["email"],
			"password":           data["password"],
			"token":              data["token"],
			"tel":                data["tel"],
			"raison_sociale":     data["raison_sociale"],
			"creation_du_compte": time.Now().Format("2006-01-02 15:04:05"),
			"pts_fidelite":       data["pts_fidelite"],
			"idADRESSE":          lastAddressID,
			"image":              data["image"],
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
			"lastCompanyInserted": lastInsertID,
		}, http.StatusOK)
	}
}
