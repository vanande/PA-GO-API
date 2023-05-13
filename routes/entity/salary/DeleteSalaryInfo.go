package salary

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func DeleteSalaryInfo(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "DELETE":
		data := libraries.Body(w, req)

		err := query.DeleteQuery("participant_info", "idParticipant = ? AND idClient = ? AND idInfo = ?", data["idp"], data["idc"], data["idi"])
		if err != nil {
			fmt.Println(w, err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully deleted data",
		}, http.StatusOK)
	}
}
