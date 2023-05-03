package salary

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func AddSalaryInfo(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		_, err := query.InsertQuery("participant_info", map[string]interface{}{
			"idParticipant": data["idp"],
			"idClient":      data["idc"],
			"idInfo":        data["idi"],
		})

		if err != nil {
			fmt.Errorf("Error: %v", err)

			libraries.Response(w, map[string]interface{}{
				"message": "Duplicate entry",
			}, http.StatusOK)
			return
		}

	}
	libraries.Response(w, map[string]interface{}{
		"message": "Successfully added",
	}, http.StatusOK)

}
