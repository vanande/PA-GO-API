package main

import (
	"TogetherAndStronger/routes/db/query"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	data := map[string]interface{}{
		"name":  "Johnny",
		"email": "johnny@example.com",
		"age":   25,
		"city":  "Paris",
	}

	query.InsertQuery("users", data)
}
