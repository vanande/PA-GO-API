package salary

import (
	"TogetherAndStronger/libraries"
	dbinit "TogetherAndStronger/routes/db/init"
	"fmt"
	"html"
	"net/http"
)

type Info struct {
	ID   string
	Name string
}

// Request : SELECT infos.nom FROM participant_info JOIN infos ON infos.idINFO = participant_info.idINFO WHERE idPARTICIPANT = 1 AND idCLIENT = 1;
func GetSalaryInfo(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		query := fmt.Sprintf("SELECT infos.idINFO, infos.nom FROM participant_info JOIN infos ON infos.idINFO = participant_info.idINFO WHERE idPARTICIPANT = %s AND idCLIENT = %s", data["idp"], data["idc"])
		fmt.Println(html.EscapeString(query))

		db, err := dbinit.InitDB()
		if err != nil {
			fmt.Errorf("Database connection failed : %v ", err)
		}

		defer db.Close()

		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		}

		var listInfo []Info

		for rows.Next() {
			var i Info
			err := rows.Scan(&i.ID, &i.Name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(i)
			listInfo = append(listInfo, i)
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    listInfo,
		}, http.StatusOK)
	}
}
