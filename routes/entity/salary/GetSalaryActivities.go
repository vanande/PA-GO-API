package salary

import (
	"TogetherAndStronger/libraries"
	dbinit "TogetherAndStronger/routes/db/init"
	"fmt"
	"html"
	"net/http"
)

type ListActivite struct {
	ID            string
	Nom           string
	Description   string
	NbPersonneMin string
	NbPersonneMax string
	PrixMin       string
	PrixMax       string
    HeureDebut string
    HeureFin string
    DateDebut string
    DateFin string
	Image         string
}

func GetSalaryActivities(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		data := libraries.Body(w, req)

		if data["idp"] == nil || data["idc"] == nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Empty parameters",
			}, http.StatusOK)
			return
		}
		if data["idp"] == "" || data["idc"] == "" {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusOK)
			return
		}
		query := fmt.Sprintf(`
			SELECT 
				list_activite.*, 
				teambuilding_activite.heure_debut, 
				teambuilding_activite.heure_fin, 
				teambuilding_activite.date_debut, 
				teambuilding_activite.date_fin 
			FROM 
				participant_equipe 
				JOIN teambuilding_activite ON participant_equipe.idTEAM_BUILDING = teambuilding_activite.idTEAM_BUILDING 
				JOIN activite ON teambuilding_activite.idActivite = activite.idActivite 
				JOIN list_activite ON activite.idlist_activite = list_activite.idlist_activite 
			WHERE 
				participant_equipe.idPARTICIPANT = %s 
				AND participant_equipe.idCLIENT = %s`,
			data["idp"], data["idc"])
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

		var listActivites []ListActivite

		for rows.Next() {
			var la ListActivite
			err := rows.Scan(&la.ID, &la.Nom, &la.Description, &la.NbPersonneMin, &la.NbPersonneMax, &la.PrixMin, &la.PrixMax, &la.Image, &la.HeureDebut, &la.HeureFin, &la.DateDebut, &la.DateFin)
			if err != nil {
				fmt.Println(err)
			}
			la.Image = fmt.Sprintf("https://%s/public/activity/%s", req.Host, la.Image)
			listActivites = append(listActivites, la)
		}
		fmt.Println(listActivites)

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    listActivites,
		}, http.StatusOK)
	}
}
