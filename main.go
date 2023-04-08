package main

import (
	db "TogetherAndStronger/routes/db/init"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	db.Main()

	/*
		// Start the server
		http.HandleFunc("/faq", faq.Faq)
		http.ListenAndServe(":9000", nil)
	*/
}
