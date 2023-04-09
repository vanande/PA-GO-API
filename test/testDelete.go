package main

import (
	"TogetherAndStronger/routes/db/query"
	_ "strings"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	query.DeleteQuery("users", "id = ?", 4)
}
