package room

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type Room struct {
	IdSalle       int    `json:"idPRESTATAIRE"`
	Nom           string `json:"nom"`
	NumSalle      string `json:"num_salle"`
	Prix          int    `json:"prix"`
	Description   string `json:"description"`
	Disponibilite int    `json:"disponibilite"`
<<<<<<< HEAD
	idLIEU        string `json:"idl"`
=======
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
	Image         string `json:"image"`
}

func GetRoom(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		if data["id"] == nil || data["id"] == "" {
			libraries.Response(w, map[string]interface{}{
				"message": "Bad parameters",
			}, http.StatusBadRequest)
			return
		}

<<<<<<< HEAD
		rows, err := query.SelectQuery("salle", []string{"*"}, map[string]interface{}{"idSALLE": data["id"]})
=======
		rows, err := query.SelectQuery("lieu", []string{"*"}, map[string]interface{}{"idLIEU": data["id"]})
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var room Room

		for rows.Next() {
<<<<<<< HEAD
			err := rows.Scan(&room.IdSalle, &room.Nom, &room.NumSalle, &room.Prix, &room.Description, &room.Disponibilite, &room.idLIEU, &room.Image)
=======
			err := rows.Scan(&room.IdSalle, &room.Nom, &room.NumSalle, &room.Prix, &room.Description, &room.Disponibilite, &room.Image)
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
			if err != nil {
				fmt.Println(err)
			}
			room.Image = fmt.Sprintf("https://%s/public/activity/%s", req.Host, room.Image)
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    room,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
