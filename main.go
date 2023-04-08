package main

import (
	"TogetherAndStronger/routes/db"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	data := map[string]string{
		"name":  "John",
		"email": "john@example.com",
	}

	db.InsertQuery("users", data)

	/*
		// Start the server
		http.HandleFunc("/faq", faq.Faq)
		http.ListenAndServe(":9000", nil)
	*/
}
