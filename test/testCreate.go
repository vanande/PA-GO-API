package main

import (
	"TogetherAndStronger/routes/db"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	data := map[string]string{
		"name":  "Johnny",
		"email": "johnny@example.com",
	}

	db.InsertQuery("users", data)
}
