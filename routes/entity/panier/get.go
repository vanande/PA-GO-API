package panier

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

type S struct {
	IdPanier string `json:"idp"`
	IdClient string `json:"idc"`
	Statut   string `json:"statut"`
	Date     string `json:"date"`
}

func Get(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)

		id, OK := data["id"].(string)
		if !OK {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid parameters",
			}, http.StatusBadRequest)
			return
		}

		// @TODO : mysql> select a.idActivite, la.idlist_activite, o.list_option FROM panier_activite p INNER JOIN activite a ON a.idActivi
		//te = p.idACTIVITE INNER JOIN list_activite la ON a.idlist_activite = la.idlist_activite INNER JOIN tas.option o ON  WHER
		//E p.idPANIER = 1;
		//ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'WHERE p.idPANIER = 1' at line 1
		//mysql> select a.idActivite, la.idlist_activite, o.list_option FROM panier_activite p INNER JOIN activite a ON a.idActivite = p.idACTIVITE INNER JOIN list_activite la ON a.idlist_activite = la.idlist_activite INNER JOIN tas.option o ON o.idActivite = a.idActivite  WHERE p.idPANIER = 1;
		//ERROR 1054 (42S22): Unknown column 'o.list_option' in 'field list'
		//mysql> select a.idActivite, la.idlist_activite, o.list_option FROM panier_activite p INNER JOIN activite a ON a.idActivite = p.idACTIVITE INNER JOIN list_activite la ON a.idlist_activite = la.idlist_activite INNER JOIN tas.option o ON o.idA
		//ctivite = a.idActivite  WHERE p.idPANIER = 1;
		//ERROR 1054 (42S22): Unknown column 'o.list_option' in 'field list'
		//mysql> select a.idActivite, la.idlist_activite, o.list_option FROM panier_activite p INNER JOIN activite a ON a.idActivite = p.idACTIVITE INNER JOIN list_activite la ON a.idlist_activite = la.idlist_activite INNER JOIN tas.option o ON o.idA
		//ctivite = a.idActivite  WHERE p.idPANIER = 1;


		//						-->	change here <--  									-->	and here <--
		rows, err := query.SelectQuery("panier", []string{"*"}, map[string]interface{}{"idPANIER": id})
		if err != nil {
			fmt.Errorf("select failed : %v", err)
		}

		var s S

		for rows.Next() {
			// let copilot do the job -> ctrl+x (the line under 'err := ...') -> wait a bit -> tab
			err := rows.Scan(&s.IdPanier, &s.IdClient, &s.Statut, &s.Date)
			if err != nil {
				fmt.Println(err)
			}
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    s,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Method not allowed",
		}, http.StatusMethodNotAllowed)
	}
}
