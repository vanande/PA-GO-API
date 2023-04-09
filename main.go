package main

import (
	"TogetherAndStronger/routes/db/api"
	"TogetherAndStronger/routes/faq"
	"net/http"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	//db.Main()

	// Start the server
	http.HandleFunc("/faq", faq.Faq)
	http.HandleFunc("/db/create", api.Create)
	http.HandleFunc("/db/select", api.Select)
	http.HandleFunc("/db/update", api.Update)
	http.HandleFunc("/db/delete", api.Delete)
	http.ListenAndServe(":9000", nil)
}
