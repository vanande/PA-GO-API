package list

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Category struct {
	IdCategory   int    `json:"idCategory"`
	CategoryNom  string `json:"nom"`
	CategoryDesc string `json:"description"`
}

type Option struct {
	IdOption   int     `json:"idOption"`
	OptionPrix float64 `json:"prix"`
	OptionNom  string  `json:"nom"`
	OptionDesc string  `json:"description"`
}

type Activite struct {
	IdActivite       string     `json:"id"`
	Nom              string     `json:"nom"`
	Description      string     `json:"description"`
	NbPersonneMin    int        `json:"nb_personne_min"`
	NbPersonneMax    int        `json:"nb_personne_max"`
	PrixMin          float64    `json:"prix_min"`
	PrixMax          float64    `json:"prix_max"`
	Image            string     `json:"image"`
	Note             float64    `json:"note"`
	CategoryActivite []Category `json:"category_activite"`
	Options          []Option   `json:"options"`
}

func ListActivities(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":
		enableCors(&w)

		// tables in the right order
		tables := []string{"list_activite la", "list_option o", "category_activite ca", "category c"}
		// columns to look for
		columns := []string{"la.idlist_activite", "la.nom", "la.description", "la.nb_personne_min", "la.nb_personne_max", "la.prix_min", "la.prix_max", "la.image", "la.note", "c.idCategory", "c.nom", "c.description", "o.idlist_option", "o.prix", "o.nom", "o.description"}
		// on what to join
		joins := []map[string]string{
			{"la.idlist_activite": "o.idlist_activite"},
			{"la.idlist_activite": "ca.idlist_activite"},
			{"ca.idCategory": "c.idCategory"},
		}
		// the where at the end
		conditions := map[string]interface{}{}

		rows, err := query.SelectWithInnerJoin(tables, columns, joins, conditions)
		if err != nil {
			fmt.Println(err)
			libraries.Response(w, map[string]interface{}{
				"message": "No data found",
			}, http.StatusNotFound)
			return
		}

		var res []Activite
		var lastA Activite
		var lastO Option
		var lastC Category

		for rows.Next() {
			var a Activite
			var o Option
			var c Category

			err = rows.Scan(&a.IdActivite, &a.Nom, &a.Description, &a.NbPersonneMin, &a.NbPersonneMax, &a.PrixMin, &a.PrixMax, &a.Image, &a.Note, &c.IdCategory, &c.CategoryNom, &c.CategoryDesc, &o.IdOption, &o.OptionPrix, &o.OptionNom, &o.OptionDesc)
			if err != nil {
				fmt.Println(err)
			}

			if lastA.IdActivite != a.IdActivite {
				a.CategoryActivite = append(a.CategoryActivite, c)
				a.Options = append(a.Options, o)
				res = append(res, a)
				lastA = a
				lastC = c
				lastO = o
			} else {
				if lastC.IdCategory != c.IdCategory {
					res[len(res)-1].CategoryActivite = append(res[len(res)-1].CategoryActivite, c)
					lastC = c
				}
				if lastO.IdOption != o.IdOption {
					res[len(res)-1].Options = append(res[len(res)-1].Options, o)
					lastO = o
				}
			}
		}

		libraries.Response(w, map[string]interface{}{
			"message": "Successfully fetched data",
			"data":    res,
		}, http.StatusOK)

	case "OPTIONS":
		fmt.Println("Preflight handled")
		w.WriteHeader(http.StatusOK)
	}
}
