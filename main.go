package main

import (
	"TogetherAndStronger/routes/faq"
	"net/http"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	// Start the server
	http.HandleFunc("/faq", faq.Faq)
	http.ListenAndServe(":9000", nil)

}
