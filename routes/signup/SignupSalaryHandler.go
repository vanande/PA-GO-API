package signup

import (
	"TogetherAndStronger/libraries"
	"TogetherAndStronger/routes/db/query"
	"net/http"
	"strconv"
)

// The request to build is the following :
//INSERT INTO `participant`(
// `nom`, `prenom`, `telephone`, `age`, `sexe`, `idCLIENT`, `idEQUIPE`)
//VALUES ('[value-1]','[value-2]','[value-3]','[value-4]','[value-5]','[value-6]','[value-7]')

func SignupSalary(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		data := libraries.Body(w, req)
		if data == nil {
			return
		}

		if data["nom"] == nil || data["prenom"] == nil || data["telephone"] == nil || data["age"] == nil || data["sexe"] == nil || data["idCLIENT"] == nil || data["idEQUIPE"] == nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Missing data",
			}, http.StatusBadRequest)
			return
		}

		if !libraries.Correct(data["nom"].(string), "nom") ||
			!libraries.Correct(data["prenom"].(string), "prenom") ||
			!libraries.Correct(data["telephone"].(string), "telephone") ||
			!libraries.Correct(data["age"].(string), "age") ||
			!libraries.Correct(data["sexe"].(string), "sexe") ||
			!libraries.Correct(data["idCLIENT"].(string), "idCLIENT") ||
			!libraries.Correct(data["idEQUIPE"].(string), "idEQUIPE") {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusBadRequest)
			return
		}

		lastInsertId, err := query.InsertQuery("participant", map[string]interface{}{
			"nom":       data["nom"],
			"prenom":    data["prenom"],
			"telephone": data["telephone"],
			"age":       data["age"],
			"sexe":      data["sexe"],
			"idCLIENT":  data["idCLIENT"],
		})
		if err != nil {
			libraries.Response(w, map[string]interface{}{
				"message": "Invalid data",
			}, http.StatusBadRequest)
			return
		}

		idc, _ := strconv.Atoi(data["idCLIENT"].(string))
		ide, _ := strconv.Atoi(data["idEQUIPE"].(string))

		token := lastInsertId*100 + idc*10 + ide

		_, err = query.InsertQuery("participant", map[string]interface{}{
			"token": token,
		})

		libraries.Response(w, map[string]interface{}{
			"message": "Salary " + data["nom"].(string) + " " + data["prenom"].(string) + " successfully signed up",
			"id":      lastInsertId,
			"token":   token,
		}, http.StatusOK)

	default:
		libraries.Response(w, map[string]interface{}{
			"message": "Invalid request",
		}, http.StatusBadRequest)
	}
}
