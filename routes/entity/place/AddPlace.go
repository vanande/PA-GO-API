package place

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

// AddPlace create a new place and the associated address in the lieu and adresse tables in the database.
func AddPlace(w http.ResponseWriter, req *http.Request) {
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

		lastPlaceID, err := query.InsertQuery("lieu", map[string]interface{}{
			"nom":       data["nom"],
			"idADRESSE": lastAddressID,
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
			"lastPlaceInserted":   lastPlaceID,
		}, http.StatusOK)
	}
}
