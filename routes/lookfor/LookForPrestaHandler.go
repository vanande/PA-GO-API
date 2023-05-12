package lookfor

import (
	"TogetherAndStronger/libraries"
	"net/http"
)

func LookForPresta(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		libraries.Response(w, map[string]interface{}{
			"message": "Obsolete, use /presta/getPresta instead",
		}, http.StatusMovedPermanently)
	}
}
