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
	// go run ..\..\go\go1.20.1\src\crypto\tls\generate_cert.go -host="127.0.0.1"
	http.ListenAndServeTLS(":9000", "cert.pem", "key.pem", nil)
}
